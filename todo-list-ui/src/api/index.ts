import axios from "axios";

export const api = axios.create({
  baseURL: "http://fedora:8080/api/v1",
  withCredentials: false
})
