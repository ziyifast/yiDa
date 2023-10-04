import {get, post, postJson} from "./http"

export const addUser = (params) => post(`/user/add`, params);

export const getUserById = (id)=> get(`/user/getUserById/${id}`)

export const getUserByName = (name)=> get(`/user/getUserByName/${name}`)

// export const loginIn = (username, password) => post(`/user/login/${username}/${password}`);
export const loginIn = (params) => post(`/user/login`, params);

export const getData = () => get('/data');

// export const testJson = (params) => postJson(`/user/testJson`, params)


//=======activity_detail相关========
export const getActivityList = () => get('/activity_detail/list');

export const updateActitityDetail = (params) => postJson(`/activity_detail/update`, params);