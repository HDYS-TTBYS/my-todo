/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

/**
 * Todoオブジェクトの共通スキーマ
 */
export type ToDo = {
    /**
     * ユニークID
     */
    readonly id: number;
    /**
     * Todoのタイトル
     */
    title: string;
    /**
     * Todoの説明
     */
    description?: string;
    /**
     * Todoの完了・未完了を示すフラグ（trueが完了）
     */
    is_complete?: boolean;
    /**
     * Todoを担当する人の名前
     */
    assagin_person?: string | null;
    /**
     * Todoが作成された時刻（UNIXタイムで単位は秒）
     */
    readonly created_at: number;
    /**
     * Todoが更新された時刻（UNIXタイムで単位は秒）
     */
    readonly updated_at?: number;
};

