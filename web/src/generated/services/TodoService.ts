/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Error } from '../models/Error';
import type { ToDo } from '../models/ToDo';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class TodoService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Get Todo Info by Todo ID
     * todo IDが一致するtodoの情報を取得
     * @param id 既存のtodoのID
     * @returns ToDo Todo Found
     * @throws ApiError
     */
    public getTodoId(
        id: number,
    ): CancelablePromise<ToDo> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/api/todo/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                404: `Todo Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Delete Todo Info by Todo ID
     * todo IDが一致するtodoの情報を削除
     * @param id 既存のtodoのID
     * @returns Error Todo Deleted
     * @throws ApiError
     */
    public deleteTodoId(
        id: number,
    ): CancelablePromise<Error> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/api/todo/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                404: `Todo Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update Todo Info by Todo ID
     * todo IDが一致するtodoの情報を更新
     * @param id 既存のtodoのID
     * @param requestBody APIに必要なフィールドを投稿して、新しいToDoを作成します
     * @returns ToDo Todo Patched
     * @throws ApiError
     */
    public updateTodoId(
        id: number,
        requestBody?: {
            /**
             * Todoのタイトル
             */
            title: string;
            description?: string;
            assigin_person: string;
            /**
             * Todoの完了・未完了を示すフラグ（trueが完了）
             */
            is_complete: boolean;
        },
    ): CancelablePromise<ToDo> {
        return this.httpRequest.request({
            method: 'PATCH',
            url: '/api/todo/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                404: `Todo Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create New Todo
     * Todoを新規で作成する
     * @param requestBody APIに必要なフィールドを投稿して、新しいToDoを作成します
     * @returns ToDo Todo Created
     * @throws ApiError
     */
    public postTodo(
        requestBody?: {
            /**
             * Todoのタイトル
             */
            title: string;
            description?: string;
            assigin_person: string;
        },
    ): CancelablePromise<ToDo> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/api/todo',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Missing Required Information(Bad Request)`,
                409: `Todo Already Exits`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Get All Todo
     * すべてのToDoを created_at descで取得する
     * ただし、ToDoの総数が20を超えると、それ以上のToDoは返却されない
     * @param offset 結果セットの収集を開始する前にスキップするアイテムの数
     * @returns any OK
     * @throws ApiError
     */
    public getTodos(
        offset: number,
    ): CancelablePromise<{
        total: number;
        todos?: Array<ToDo>;
    }> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/api/todos',
            query: {
                'offset': offset,
            },
            errors: {
                400: `Bad Request`,
                404: `Todo Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

}
