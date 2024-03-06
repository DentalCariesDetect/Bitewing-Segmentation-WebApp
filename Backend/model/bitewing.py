from typing import Optional,TypeVar
from pydantic import BaseModel



T = TypeVar('T')

class BitewingModel(BaseModel):
    id: Optional[int]= None
    image_id: Optional[int]= None
    name: Optional[str]= None
    side: Optional[str]= None
    status: Optional[str]= None

class BitewingInsert(BaseModel):
    image_id: int
    name: Optional[str] = None
    side: Optional[str]= None
    status: str
    created_at: Optional[str] = None