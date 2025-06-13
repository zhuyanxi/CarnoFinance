// place files you want to import through the `$lib` alias in this folder.

export interface ErrMsg {
    code: number;
    message: string;
    hidden: boolean;
}

export function getFormattedCurrentDate(): string {
    const today = new Date();
    const year = today.getFullYear();
    const month = (today.getMonth() + 1).toString().padStart(2, '0');
    const day = today.getDate().toString().padStart(2, '0');
    return `${year}-${month}-${day}`;
}