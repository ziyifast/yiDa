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

export const getActivityListByPage = (startPage, pageSize) => get(`/activity_detail/list/${startPage}/${pageSize}`);

export const updateActitityDetail = (params) => postJson(`/activity_detail/update`, params);

export const getActivityListByUid = (uid) => get(`/activity_detail/get/${uid}`);

export const addActivityDetail = (params) => post(`/activity_detail/add`, params);


//=======user_activity相关========
export const deleteUserActivityByUidAndAid = (uid, aid) => post(`/user_activity/delete/${uid}/${aid}`)

export const delUserActivity = (uid, aid) => post(`/user_activity/delete/${uid}/${aid}`)

export const addUserActivity = (uid, aid) => post(`/user_activity/add/${uid}/${aid}`);  