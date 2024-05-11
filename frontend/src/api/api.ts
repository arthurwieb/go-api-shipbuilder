import axios from "axios";
import { ShipModel } from "../models/shipModel";

const URL = "http://localhost:3000/";

async function getShips(): Promise<ShipModel[]> {
  const response = await axios.get<ShipModel[]>(`${URL}ships`);
  return response.data;
}

async function getShipTypes(): Promise<ShipModel[]> {
  const response = await axios.get<ShipModel[]>(`${URL}shiptypes`);
  return response.data;
}
// async function updateUserName(userId: string, name: string): Promise<User> {
//   const response = await axios.put<User>(`${URL}users/${userId}`, { name });

//   return response.data;
// }

export const api = {
  getShips,
  getShipTypes,
};
