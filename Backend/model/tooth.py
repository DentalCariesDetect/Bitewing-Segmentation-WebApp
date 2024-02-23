from typing import Optional,TypeVar
from fastapi import File, UploadFile
from pydantic import BaseModel

T = TypeVar('T')

class ToothModel(BaseModel):
    id: Optional[int]
    result_id: Optional[int]
    image_id: Optional[int]
    name: Optional[str]
    type_tooth: Optional[str]
    type_caries: Optional[str]
    numbering: Optional[int]
    filing: Optional[str]
    position: Optional[str]
    status: Optional[str]

class ToothInsert(BaseModel):
    result_id: int
    image_id: int
    numbering: int
    position: str
    created_at: Optional[str] = None