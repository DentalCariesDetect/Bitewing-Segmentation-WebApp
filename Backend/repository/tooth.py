from model.tooth import ToothInsert
from config.database import postgres

class ToothRepository:
    @staticmethod
    async def create(data: ToothInsert):
        query = "INSERT INTO tooths (result_id, image_id,numbering,position,status,created_at) VALUES (:result_id, :image_id,:numbering,:position, :status,now()) RETURNING id"
        values = {
            "result_id": data.result_id,
            "image_id": data.image_id,
            "numbering": data.numbering,
            "position": data.position,
            "status": "Active"
        }
        result = await postgres.database.execute(query=query, values=values)
        return result