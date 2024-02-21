from fastapi import APIRouter
from model.response import ResponseModel
router = APIRouter(
    prefix="/api/segmentation",
    tags=["segmentation"],
    responses={404: {"description": "Not found"}},
)

@router.post("/api/crop",response_model= ResponseModel,response_model_exclude_none=True)
async def crop(file: UploadFile = File(...)):
    image = cropping.save_image(file)
    data_path = cropping.cropping(image)
    image_path = data_path[0]
    image_path_crop = data_path[1]
        
    images_data_base64 = []
    for path in image_path:
        img = Image.open(path)
        images_data_base64.append(image_to_base64(img))

        image_data_crop = Image.open(image_path_crop)
        image_data_crop_base64 = image_to_base64(image_data_crop)
        
        response = dict(
            crop_img=image_data_crop_base64,
            list_crop_img=images_data_base64
        )   
    return response