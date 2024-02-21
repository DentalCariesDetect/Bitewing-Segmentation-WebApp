from contextlib import asynccontextmanager
from fastapi.middleware.cors import CORSMiddleware 
from fastapi import FastAPI
from service.ImageService import image_to_base64
import uvicorn
import os
import cropping_model as cropping
from PIL import Image

def initialize_app():

    @asynccontextmanager
    async def lifespan(app:FastAPI):
        print("Starting app")
        #connect to database
        yield 
        print("Stopping app")
        
        
    app =FastAPI(lifespan=lifespan)

    app.add_middleware(
        CORSMiddleware,
        allow_origins=["http://localhost:3000"],  # Replace with the actual origin of your web app
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

    @app.get("/")
    def read_root():
        #get env variable
        return {"Hello": os.environ.get("NODE_ENV")}




    return app

app = initialize_app()
    
if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=80)