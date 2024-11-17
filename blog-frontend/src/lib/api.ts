import { API_BASE_URL } from "@/utils/constants";
import axios from "axios";

const api = axios.create({
    baseURL: API_BASE_URL,
    timeout: 1000,
    headers:{
        "Content-Type": "application/json"
    }
})

export default api