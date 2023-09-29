import Axios from "axios";
import {get, post} from "./http"
//=======用户相关========
export const addUser = () => post('/user/addUser');

export const loginIn = (username, password) => post(`/user/login/${username}/${password}`);

export const getData = () => get('/data');