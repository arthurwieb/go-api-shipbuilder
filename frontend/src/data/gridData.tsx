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
