from model.bitewing import BitewingInsert
from config.database import postgres

class BitewingRepository:
    @staticmethod
    async def create(data: BitewingInsert):
        query = "INSERT INTO bitewings (image_id, status,created_at) VALUES (:image_id, :status,now()) RETURNING id"
        values = {
            "image_id": data.image_id,
            "status": data.status
        }
        result = await postgres.database.execute(query=query, values=values)
        return result