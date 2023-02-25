## SPDX-License-Identifier: Apache-2.0

from fastapi import FastAPI

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}
