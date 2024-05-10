import "./ship.scss"
import React, { useState } from 'react';
import DataTable from "../../components/dataTable/DataTable";
import { myShipColumns } from '../../data/gridData';
import { useQuery } from "react-query";
import { api } from "../../api/api"
import Add from "../../components/add/AddShip";
import { shipColumns } from "../../data";

const Ship: React.FC = () => {

  const { data, isError, isLoading } = useQuery('ship-list', api.getShips);

  const [open, setOpen] = useState(false)

  return (
    <div className="ship">
      <div className="info">
        <h1>My Ships</h1>
        <button onClick={() => setOpen(true)}>Add New Ship</button>
      </div>
      {isLoading && <div>Loading...</div>}
      {isError && <div>Ocorreu algum erro... F</div>}
      {data && 
        <DataTable slug="ship" columns={myShipColumns} rows={data} />
      }
      {open && <Add slug="ship" columns={shipColumns} setOpen={setOpen} />}
    </div>
  )
};

export default Ship;
