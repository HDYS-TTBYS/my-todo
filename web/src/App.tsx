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
    path: "/todo",
    element: <Index />,
    errorElement: <ErrorPage />
  },
  {
    path: "/todo/create",
    element: <Create />,
  },
  {
    path: "/todo/delete/:id",
    element: <Delete />,
  },
  {
    path: "/todo/detail/:id",
    element: <Detail />,
  },
  {
    path: "/todo/update/:id",
    element: <Update />,
  },
]);

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      learn react
      <RouterProvider router={router} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>

  );
}

export default App;
