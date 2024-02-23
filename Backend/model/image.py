from typing import Optional,TypeVar
from pydantic import BaseModel



T = TypeVar('T')

class ImageModel(BaseModel):
    size_x: Optional[int]= None
    size_y: Optional[int]= None
    path: Optional[str]= None
    status: Optional[str]  = None

class ImageInsert(BaseModel):
    size_x: Optional[int] = None
    size_y: Optional[int] = None
    path: str
    status: str
    