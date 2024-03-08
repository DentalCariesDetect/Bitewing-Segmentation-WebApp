import time
from fastapi import APIRouter, File, Form, UploadFile
from fastapi import Depends, HTTPException, status
from fastapi.security import OAuth2PasswordBearer

from model.response import ResponseModel
from service.cropping import CroppingService
from service.auth import AuthService
from service.segmentation import SegmentationService
from model.segmentation import SegmentationRequest
from PIL import Image

import base64
from io import BytesIO
from typing import List

def image_to_base64(image_path: str) -> str:
    """Convert an image to a base64 string."""
    with Image.open(image_path) as image:
        buffered = BytesIO()
        image.save(buffered, format="JPEG")  # Adjust the format based on your image format
        img_str = base64.b64encode(buffered.getvalue()).decode("utf-8")
    return img_str

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")

router = APIRouter(
    prefix="/api/segmentation",
    tags=["segmentation"],
    responses={404: {"description": "Not found"}},
)
@router.post("/crop",response_model= ResponseModel,response_model_exclude_none=True)
async def cropping(file: UploadFile = File(...),patient_id:int=Form(...) ,token:str = Depends(oauth2_scheme)):
    start_time = time.time()
    dentist_id = await AuthService.verifyAuth(token)
    image = CroppingService.save_image(file)
    data_path = CroppingService.cropping(image)
    bitewing_image_path = data_path[1]
    image_path_crop = data_path[0]
    #add to database
    segmentation_request = SegmentationRequest(
        dentist_id=dentist_id,
        patient_id=patient_id,
        bitewing_path=bitewing_image_path,
        list_crop_img=image_path_crop,
    )
    result = await SegmentationService.insert(segmentation_request)
    if result == None:
        raise HTTPException(
            status_code=400,
            detail="Failed to insert data",
        )

    end_time = time.time()
    print(f"Total time: {end_time - start_time} seconds")
    return ResponseModel(data=result, message="Image has been cropped successfully")


@router.post("/test_crop", response_model=ResponseModel, response_model_exclude_none=True)
async def crop(file: UploadFile = File(...)):
    start_time = time.time()  # Start timing

    # Save the uploaded image and get its path
    image = CroppingService.save_image(file)

    # Process the image to get cropped parts and overview image
    cropped_images_data, overview_image_path = CroppingService.cropping(image)

    # Encode the overview image to base64
    overview_image_base64 = image_to_base64(overview_image_path)

    # Encode all cropped images to base64
    cropped_images_base64 = [{
        "position": data["position"],
        "numbering": data["numbering"],
        "base64_image": image_to_base64(data["image_path"])  # Encode each cropped image
    } for data in cropped_images_data]

    response = dict(
        crop_img=overview_image_base64,  # Overview image as base64
        list_crop_img=cropped_images_base64  # List of cropped images as base64
    )

    end_time = time.time()
  
    print(f"Total time: {end_time - start_time} seconds")

    return ResponseModel(data=response, message="Image has been cropped successfully")