from model.image import ImageInsert
from config.database import postgres

class ImageRepository:
    @staticmethod
    async def create(data: ImageInsert):
        query = "INSERT INTO images (path, status,created_at) VALUES (:path, :status,now()) RETURNING id"
        values = {
            "path": data.path,
            "status": data.status
        }
        result = await postgres.database.execute(query=query, values=values)
        return result
        