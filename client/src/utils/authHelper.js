import cookie from "js-cookie";
import { STORAGE_USER, TOKEN } from './constants';

// set in cookie
export const setCookie = (key, value) => cookie.set(key, value, {
      expires: 1,
    })

// remove from cookie
export const removeCookie = (key) => cookie.remove(key, {
      expires: 1,
    });

// get from cookie such as stored token
// will be useful when we need to make request to server with token
export const getCookie = (key) => cookie.get(key);

// set in localstorage
export const setLocalStorage = (key, value) => localStorage.setItem(key, JSON.stringify(value));

// get from localstorage
export const getFromLocalStorage = (key) => localStorage.getItem(key)

// remove from localstorage
export const removeLocalStorage = (key) => localStorage.removeItem(key);

// authenticate user by passing data to cookie and localstorage during signin
export const authenticate = (response, next) => {
  console.log("AUTHENTICATE HELPER ON SIGNIN RESPONSE", response);
  setCookie(TOKEN, response.token);
  setLocalStorage(STORAGE_USER, response.user);
  next();
};
// access user info from localstorage
export const isAuth = () => {
  if (window !== "undefined") {
    const cookieChecked = getCookie(TOKEN);
    if (cookieChecked) {
      if (localStorage.getItem(STORAGE_USER)) {
        return JSON.parse(localStorage.getItem(STORAGE_USER));
      } else {
        return false;
      }
    }
  }
};


export const signout = (next) => {
  removeCookie(TOKEN);
  removeLocalStorage(STORAGE_USER);
  next();
};

export const updateUser = (response, next) => {
  console.log("UPDATE USER IN LOCALSTORAGE HELPERS", response);
  if (typeof window !== "undefined") {
    let auth = JSON.parse(localStorage.getItem("user"));
    auth = response.data;
    localStorage.setItem(STORAGE_USER, JSON.stringify(auth));
  }
  next();
};
