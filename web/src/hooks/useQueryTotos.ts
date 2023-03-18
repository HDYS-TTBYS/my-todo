import { useQuery } from "react-query";
import { GetTodos200Response } from "../types/typescript-fetch";

export const useQueryTodos = (offset: number) => {
  const getTodos = async () => {
    const res = await fetch(`${process.env.REACT_APP_REST_URL}/api/todos?offset=${offset}`)
    if (!res.ok) {
      return res.json() as Promise<GetTodos200Response>
    }
    return res.json() as Promise<Error>
  }

  return useQuery({
    queryKey: "todos",
    queryFn: getTodos,
  })
}
