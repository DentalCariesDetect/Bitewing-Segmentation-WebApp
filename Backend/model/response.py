from typing import Optional,TypeVar
from pydantic import BaseModel

T = TypeVar('T')

class ResponseModel(BaseModel):
    data: T
    message: Optional[str] = None
    code: Optional[int] = 200
