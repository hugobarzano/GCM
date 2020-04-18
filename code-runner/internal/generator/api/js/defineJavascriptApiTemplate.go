package js

const JsApiTemplate = `
'use strict';

var DATA = new Map();

// API Operations
const listAll = (req, res) => {
    var response = [];
    for (var [, value] of DATA) {
        response.push(value)
    }
    return res.status(200).send(response);
};

const getOne = (req, res) => {
    var id = req.params.id;
    if (DATA.has(id)){
        return res.json(DATA.get(id))
    }
    return res.status(404).json({"error": "Object with id "+id+" not found"});
};

const createOne = (req, res) => {
    var obj = req.body;
    var id = obj.id;
    if (DATA.has(id)){
        return res.status(406).json({"error":"Object with id "+id+" already exists"})
    }
    DATA.set(obj.id,obj)
    return res.status(201).json(DATA.get(obj.id))
};

const updateOne = (req, res) => {
    var idToUpdate = req.params.id;
    if (DATA.has(idToUpdate)){
        var obj = req.body;
        obj.id = idToUpdate
        DATA.set(idToUpdate,obj)
        return res.json(obj)
    }
    return res.status(404).json({"error":"Object with id "+idToUpdate+" not found"})
};

const deleteOne = (req, res) => {
    var idToRemove = req.params.id;
    if (DATA.has(idToRemove)){
        DATA.delete(idToRemove)
        return res.json({"msg":"Object with id "+idToRemove+" successfully deleted"})
    }
    return res.status(406).json({"error":"Object with id "+idToRemove+" not found"})
};

module.exports = {
    listAll,
    getOne,
    createOne,
    updateOne,
    deleteOne
};
`
