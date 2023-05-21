import axios from "axios";
import {toast } from 'react-toastify';

export function getAssets(options) {
    return axios.get("/api/v1/assets", options)
    .then(response => {
        return Promise.resolve(response.data);       
    });
}
