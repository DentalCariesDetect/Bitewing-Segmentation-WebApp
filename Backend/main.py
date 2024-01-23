from fastapi import File, UploadFile
from fastapi.middleware.cors import CORSMiddleware 
from fastapi import FastAPI
from service.ImageService import image_to_base64
import uvicorn
import cropping_model as cropping
from PIL import Image

app =FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],  # Replace with the actual origin of your web app
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/")
def read_root():
    return {"Hello": "World"}



@app.post("/api/crop")
async def crop(file: UploadFile = File(...)):
    image = cropping.save_image(file)
    data_path = cropping.cropping(image)
    image_path = data_path[0]
    image_path_crop = data_path[1]
    
    images_data_base64 = []
    for path in image_path:
        print(path)
        img = Image.open(path)
        images_data_base64.append(image_to_base64(img))

    image_data_crop = Image.open(image_path_crop)
    image_data_crop_base64 = image_to_base64(image_data_crop)
    
    response = dict(
        crop_img=image_data_crop_base64,
        list_crop_img=images_data_base64
    )   
    
    return response
    
if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=80)