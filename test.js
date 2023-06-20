async function Fetch(){
const data = await fetch("http://localhost:8000/json/users/").then(response => response.json());
console.log(data);
}

Fetch();
