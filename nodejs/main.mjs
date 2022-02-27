import fetch from 'node-fetch';

// Your token
const TOKEN = "";

const date = new Date();
const today = date.getFullYear() + 
  "-" +
  String(date.getMonth() + 1).padStart(2, '0') +
  "-" +
  String(date.getDate()).padStart(2, "0");

const url = `https://api.fitbit.com/1/user/-/activities/date/${today}.json`;
const option = {
  method: "GET",
  headers: {'Authorization': `Bearer ${TOKEN}`}
};

const response = await fetch(url, option).then(r => r.json());
console.log(response);
