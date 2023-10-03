import Axios from "axios";
import {get, post} from "./http"

export const addUser = () => post('/user/addUser');

export const loginIn = (username, password) => post(`/user/login/${username}/${password}`);

export const getData = () => get('/data');


//=======活动相关========
export const getActivityList = () => get('/activity/getActivityList');