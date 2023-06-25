import axios from "axios";
import {toast } from 'react-toastify';

export function getScripts(options) {
    return axios.get("/api/v1/scripts", options)
    .then(response => {
        return Promise.resolve(response.data);       
    });
}

export function saveScript(scriptPayLoad) {
    return axios.post("/api/v1/scripts", scriptPayLoad)
    .then(data => {       
        return Promise.resolve(data);       
    });
}

export function searchResults(options) {
    return axios.get("/api/v1/results", options)
    .then(response => {
        return Promise.resolve(response.data);       
    });
}