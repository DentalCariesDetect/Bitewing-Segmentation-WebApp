from typing import Optional,TypeVar
from fastapi import File, UploadFile
from pydantic import BaseModel
from typing import List, Tuple

T = TypeVar('T')

class SegmentationModel(BaseModel):
    id: Optional[int]
    bitewing_id: Optional[int]
    dentist_id: Optional[int]
    patient_id: Optional[int]
    status: Optional[str]

class SegmentationRequest(BaseModel):
    dentist_id:int
    patient_id:int
    bitewing_path: str
    list_crop_img: List[Tuple[str, str]]



class SegmentationInsert(BaseModel):
    bitewing_id: int
    dentist_id: int
    patient_id: int
    status: str