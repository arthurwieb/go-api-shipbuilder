import "./shiptypes.scss"

import React from 'react';
import DataTable from "../../components/dataTable/DataTable";
import { shipColumns } from '../../data/gridData';
import { useQuery } from "react-query";
import { api } from "../../api/api";

const ShipType: React.FC = () => {

  const { data, isError, isLoading } = useQuery('ship-type', api.getShipTypes)

  return (
    <div className="shiptypes">
      <div className="info">
        <h1>Ship Types</h1>
      </div>
      {isLoading && <div>Loading...</div>}
      {isError && <div>Ocorreu algum erro... F</div>}
      {data &&
        <DataTable slug="shiptypes" columns={shipColumns} rows={data} />
      }
    </div>
  )
};

export default ShipType;
