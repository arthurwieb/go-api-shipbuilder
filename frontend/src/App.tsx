import {createBrowserRouter, RouterProvider} from "react-router-dom";
import { Users } from "./pages/users/Users";
import ShipType from "./pages/shiptypes/ShipType";
function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      //element: <Home/>,
      element: <ShipType/>,
    },
    {
      path: "/users",
      element: <Users/>,
    },
    {
      path: "/shiptypes",
      element: <ShipType/>,
    },
  ]);
  return <RouterProvider router = {router} />;
}

export default App
