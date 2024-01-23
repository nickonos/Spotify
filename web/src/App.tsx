import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { CallbackPage } from "./pages/Callback";
import { LoginPage } from "./pages/Login";
import { Layout } from "./components/ui/layout";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { AuthenticationProvider } from "./service/AuthenticationProvider";
import { SearchPage } from "./pages/Search";
import { HomePage } from "./pages/Home";
import { UploadPage } from "./pages/Upload";

const router = createBrowserRouter([
  {
    path: "*",
    element: <AuthenticationProvider />,
    children: [
      {
        path: "auth",
        children: [
          {
            path: "callback",
            element: <CallbackPage />,
          },
          {
            path: "login",
            element: <LoginPage />,
          },
        ],
      },
      {
        path: "*",
        element: <Layout />,
        children: [
          {
            path: "search",
            element: <SearchPage />,
          },
          {
            path: "upload",
            element: <UploadPage />
          },
          {
            path: "*",
            element: <HomePage />,
          },
        ],
      },
    ],
  },
]);

const queryClient = new QueryClient();

function App() {
  return (
    <div className="dark h-screen bg-[#0C0A09] text-white">
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </div>
  );
}

export default App;
