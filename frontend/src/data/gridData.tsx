import { GridColDef } from "@mui/x-data-grid";

export const shipColumns: GridColDef[] = [
  {
    field: "ID",
    headerName: "ID",
    width: 90,
  },
  {
    field: "Name",
    type: "string",
    headerName: "Name",
    width: 150,
  },
  {
    field: "MaxCrew",
    headerName: "Max Crew",
    width: 100,
  },
  {
    field: "Speed",
    type: "number",
    headerName: "Speed",
    width: 100,
  },
  {
    field: "Wake",
    type: "number",
    headerName: "Wake",
    width: 100,
  }
];

//Most important Stats in every ship
export const myShipColumns: GridColDef[] = [
  {
    field: "ShipTypeID",
    headerName: "Type",
    width: 90,
  },
  {
    field: "Name",
    type: "string",
    headerName: "Name",
    width: 150,
  },
  {
    field: "CannonAccBonus",
    type: "number",
    headerName: "CanACC+",
    width: 100,
  },
  {
    field: "CannonMaxDmgBonus",
    type: "number",
    headerName: "CanDMG+",
    width: 100,
  },
  {
    field: "LesserBonus",
    type: "number",
    headerName: "LCD+",
    width: 100,
  },
  {
    field: "RegularBonus",
    type: "number",
    headerName: "RCD+",
    width: 100,
  },
  {
    field: "GreaterBonus",
    type: "number",
    headerName: "GCD+",
    width: 100,
  },
  {
    field: "SpeedBonus",
    type: "number",
    headerName: "Speed+",
    width: 100,
  },
  {
    field: "DoubloonsBonus",
    type: "number",
    headerName: "Dubs+",
    width: 100,
  },
  {
    field: "FishingBonus",
    type: "number",
    headerName: "Dubs+",
    width: 100,
  },
  {
    field: "SpyBonus",
    type: "number",
    headerName: "Spy+",
    width: 100,
  },
  {
    field: "WakeBonus",
    type: "number",
    headerName: "Wake+",
    width: 100,
  }
];
