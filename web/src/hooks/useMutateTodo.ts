import { todoApi } from "../api-config";
import { useQueryClient, useMutation } from "react-query";
import { PostTodoOperationRequest, ToDo, UpdateTodoIdOperationRequest, DeleteTodoIdRequest } from "../types/typescript-fetch";

export const useMutateTodo = () => {
  const queryClient = useQueryClient()

  const useMutateCreateTodo = useMutation(
    (todo: PostTodoOperationRequest) =>
      todoApi.postTodo(todo),
    {
      onSuccess: (res: ToDo) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateUpdateTodo = useMutation(
    (todo: UpdateTodoIdOperationRequest) =>
      todoApi.updateTodoId(todo),
    {
      onSuccess: (res: ToDo) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateDeleteTodo = useMutation(
    (id: DeleteTodoIdRequest) =>
      todoApi.deleteTodoId(id),
    {
      onSuccess: (res: Error) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  return { useMutateCreateTodo, useMutateUpdateTodo, useMutateDeleteTodo }
}
