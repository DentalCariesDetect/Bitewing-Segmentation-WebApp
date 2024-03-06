from model.segmentation import SegmentationInsert
from config.database import postgres


class SegmentationRepository:
    @staticmethod
    async def create(data: SegmentationInsert):
        query = "INSERT INTO results (bitewing_id, dentist_id, patient_id, status,created_at) VALUES (:bitewing_id, :dentist_id, :patient_id, :status,now()) RETURNING id"
        values = {
            "bitewing_id": data.bitewing_id,
            "dentist_id": data.dentist_id,
            "patient_id": data.patient_id,
            "status": data.status
        }
        result = await postgres.database.execute(query=query, values=values)
        return result