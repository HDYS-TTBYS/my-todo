import { useQueryClient, useMutation } from "react-query";
import { PostTodoOperationRequest, ToDo, UpdateTodoIdOperationRequest, DeleteTodoIdRequest } from "../types/typescript-fetch";
import axios from "axios";

export const useMutateTodo = () => {
  const queryClient = useQueryClient()

  const useMutateCreateTodo = useMutation(
    (todo: PostTodoOperationRequest) =>
      axios.post(`${process.env.REACT_APP_REST_URL}/api/todo`, todo.postTodoRequest) as Promise<ToDo>,
    {
      onSuccess: (res: ToDo) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateUpdateTodo = useMutation(
    (todo: UpdateTodoIdOperationRequest) =>
      axios.patch(`${process.env.REACT_APP_REST_URL}/api/todo/${todo.id}`, todo.updateTodoIdRequest!) as Promise<ToDo>,
    {
      onSuccess: (res: ToDo) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateDeleteTodo = useMutation(
    (id: DeleteTodoIdRequest) =>
      axios.delete(`${process.env.REACT_APP_REST_URL}/api/todo/${id.id}`) as Promise<Error>,
    {
      onSuccess: (res: Error) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  return { useMutateCreateTodo, useMutateUpdateTodo, useMutateDeleteTodo }
}
