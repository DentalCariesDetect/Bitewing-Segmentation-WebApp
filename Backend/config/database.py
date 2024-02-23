
import os
from databases import Database

class DatabaseConnection:
    def __init__(self):
        username = os.getenv("POSTGRES_USER")
        password = os.getenv("POSTGRES_PASSWORD")
        host = os.getenv("POSTGRES_HOST")
        port = os.getenv("POSTGRES_PORT")
        database = os.getenv("POSTGRES_DB")
        self.url = f"postgresql://{username}:{password}@{host}:{port}/{database}"
        #self.testing_url = f"postgresql://admin:1234@localhost:5432/segmentation"
        self.database = Database(self.url)


    def getDb(self):
        return self.database
    
    async def start(self):
        await self.database.connect()
    
    async def stop(self):
        await self.database.disconnect()

postgres = DatabaseConnection()

