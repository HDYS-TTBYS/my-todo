import { useQuery } from "react-query";
import appClient from "../app-client";

export const useQueryTodos = (offset: number) => {
  const getTodos = async () => {
    return appClient.todo.getTodos(offset)
  }

  return useQuery({
    queryKey: "todos",
    queryFn: getTodos,
  })
}
