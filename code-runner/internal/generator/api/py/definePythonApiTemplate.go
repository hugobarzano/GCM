package py

const PythonApiTemplate = `
"""
Api Source Code
"""
from datetime import datetime
from flask import make_response, abort

DATA = {}

def getTimestamp():
    return datetime.now().strftime(("%Y-%m-%d %H:%M:%S"))

def listAll():
    return [DATA[key] for key in sorted(DATA.keys())]

def getOne(id):
    res = {}
    if id in DATA:
        object = DATA.get(id)
        return object
    else:
        res["error"]="Object with id {id} not found".format(id=id)
        return make_response(res, 404)

def createOne(object):
    res = {}
    id = object.get("id", None)
    if id not in DATA and id is not None:
        object["timestamp"] = getTimestamp()
        DATA[id] = object
        return make_response(DATA[id], 201)
    else:
        res["error"]="Object with id {id} already exists".format(id=id)
    return make_response(
        res, 406
    )

def updateOne(id, object):

    res = {}
    if id in DATA:
        object["id"]=id
        DATA[id]=object
        return DATA[id]
    else:
        res["error"]="Object with id {id} not found".format(id=id)
        return make_response(res, 404)

def deleteOne(id):
    res = {}
    if id in DATA:
        del DATA[id]

        res["msg"]="Object with id {id} successfully deleted".format(id=id)
        return make_response(
            res, 200
        )

    else:
        res["error"]="Object with id {id} not found".format(id=id)
        return make_response(
            res, 404
        )
`
