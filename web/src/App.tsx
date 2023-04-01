import './App.css';
import { QueryClient, QueryClientProvider } from 'react-query'
import { ReactQueryDevtools } from 'react-query/devtools'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Index from './pages/Index';
import Create from './pages/Create';
import Delete from './pages/Delete';
import Detail from './pages/Detail';
import Update from './pages/Update';
import ErrorPage from './pages/Error';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
})

const router = createBrowserRouter([
  {
    path: "/",
    element: <Index />,
    errorElement: <ErrorPage />
  },
  {
    path: "/create",
    element: <Create />,
  },
  {
    path: "/delete/:id",
    element: <Delete />,
  },
  {
    path: "/detail/:id",
    element: <Detail />,
  },
  {
    path: "/update/:id",
    element: <Update />,
  },
  {
    path: "/error",
    element: <ErrorPage />,
  },
]);

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>

  );
}

export default App;
