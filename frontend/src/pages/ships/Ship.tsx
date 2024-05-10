import "./ship.scss"

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import DataTable from "../../components/dataTable/DataTable";
import { myShipColumns } from '../../data/gridData';

interface Ship {
  ID: number;
  Name: string;
  // Add other properties as needed
}

const Ship: React.FC = () => {
  const [ship, setShip] = useState<Ship[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchShip = async () => {
      try {
        const response = await axios.get<Ship[]>('http://localhost:3000/ships');
        console.log(response)
        setShip(response.data);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching ships:', error);
        setLoading(false);
      }
    };

    fetchShip();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="ship">
      <div className="info">
        <h1>Ship Types</h1>
      </div>
      <DataTable slug="ship" columns={myShipColumns} rows={ship} />
    </div>
  )
};

export default Ship;
