import DataTable from "../../components/dataTable/DataTable"
import "./users.scss"
import { userRows, userColumns } from "../../data";
import { useState } from "react";
import Add from "../../components/add/Add";

export const Users = () => {
  const [open, setOpen] = useState(false)
  return (
    <div className="users">
      <div className="info">
        <h1>Users</h1>
        <button onClick={() => setOpen(true)}>Add New User</button>
      </div>
      <DataTable slug="users" columns={userColumns} rows={userRows} />
      {open && <Add slug="users" columns={userColumns} setOpen={setOpen} />}
    </div>
  )
}
