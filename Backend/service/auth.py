import statistics
import time
from fastapi import Depends, HTTPException
from fastapi.security import OAuth2PasswordBearer
from jose import jwt
from jose.exceptions import JWTError
import os


class AuthService:


    async def verifyAuth(token: str):
        try:
            payload = AuthService.decodeJWT(token)
            if payload is None:
                raise HTTPException(
                    status_code=401,
                    detail="Could not validate credentials",
                    headers={"WWW-Authenticate": "Bearer"},
                )
            if payload.get("exp") < int(time.time()):
                raise HTTPException(
                    status_code=401,
                    detail="Token has expired",
                    headers={"WWW-Authenticate": "Bearer"},
                )
            return payload.get("dentist_id")
        except JWTError:
            raise HTTPException(
                status_code=401,
                detail="Could not validate credentials",
                headers={"WWW-Authenticate": "Bearer"},
            )
    
    def decodeJWT(token: str):
        # secret = "miraclex77"  # ใช้ key ที่ปลอดภัยและเก็บไว้ในที่ที่ปลอดภัย
        algorithm = "HS256"
        secret = os.getenv("JWT_SECRET")

        return jwt.decode(token,secret , algorithm)