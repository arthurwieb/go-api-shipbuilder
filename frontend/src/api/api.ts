import axios from "axios";
import Ship from "../pages/ships/Ship";

const URL = "'http://localhost:3000/";

async function getShips(): Promise<Ship[]> {
  const response = await axios.get<Ship[]>(`${URL}ships`);

  return response.data;
}
// async function updateUserName(userId: string, name: string): Promise<User> {
//   const response = await axios.put<User>(`${URL}users/${userId}`, { name });

//   return response.data;
// }

export const api = {
  getShips,
};
