import { createBrowserRouter, Outlet, RouterProvider } from "react-router-dom";
import { Users } from "./pages/users/Users";
import ShipType from "./pages/shiptypes/ShipType";
import Menu from "./components/menu/Menu";
import Footer from "./components/footer/Footer";
import { Home } from "./pages/home/Home";
import "./styles/global.scss"
import Navbar from "./components/navbar/Navbar";
import Ship from "./pages/ships/Ship";
import { QueryClient, QueryClientProvider } from "react-query";

const queryClient = new QueryClient();

function App() {

  const Layout = () => {
    return (
      <div className="main">
        <Navbar />
        <div className="container">
          <div className="menuContainer">
            <Menu />
          </div>
          <div className="contentContainer">
            <QueryClientProvider client={queryClient}>
              <Outlet />
            </QueryClientProvider>
          </div>
        </div>
        <Footer />
      </div>
    )
  }
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Layout />,
      children: [
        {
          path: "/",
          element: <Home />,
        },
        {
          path: "/users",
          element: <Users />,
        },
        {
          path: "/shiptypes",
          element: <ShipType />,
        },
        {
          path: "/ships",
          element: <Ship />,
        },
      ]
    },
  ]);
  return <RouterProvider router={router} />;
}

export default App
