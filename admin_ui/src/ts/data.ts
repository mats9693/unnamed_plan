export interface User {
  userID: string;
  userName: string;
  nickname: string;

  isLocked: boolean;
  isLockedDisplay: string;

  permission: number;
  createdBy: string;
}

export function displayUserIsLocked(isLocked: boolean): string {
  return isLocked ? "已锁定" : "未锁定";
}

export interface CloudFile {
  fileID: string;
  fileName: string;
  fileURL: string;

  isPublic: boolean;
  isPublicDisplay: string;

  updateTime: number;
  updateTimeDisplay: string;

  createdTime: number;
  createdTimeDisplay: string;
}

const cloudFileURLOrigin = "https://mats9693.cn/cloud-file/";

export function generateCloudFileURL(url: string): string {
  return cloudFileURLOrigin + url
}

export function displayCloudFileIsPublic(isPublic: boolean): string {
  return isPublic ? "公开" : "非公开";
}

export function displayCloudFileTime(time: number): string {
  return (new Date(time*1000)).toString();
}
