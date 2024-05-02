import "./shiptypes.scss"

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import DataTable from "../../components/dataTable/DataTable";
import { shipColumns } from '../../data/gridData';

interface ShipType {
  ID: number;
  Name: string;
  // Add other properties as needed
}

const ShipType: React.FC = () => {
  const [shiptypes, setShipTypes] = useState<ShipType[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchShipTypes = async () => {
      try {
        const response = await axios.get<ShipType[]>('http://localhost:3000/shiptypes');
        console.log(response)
        setShipTypes(response.data);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching ship types:', error);
        setLoading(false);
      }
    };

    fetchShipTypes();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="shiptypes">
      <div className="info">
        <h1>Ship Types</h1>
      </div>
      <DataTable slug="shiptypes" columns={shipColumns} rows={shiptypes} />
    </div>
  )
};

export default ShipType;
