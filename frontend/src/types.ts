export interface MusicFile {
  uid: string;
  name: string;
  size: number;
  metadata: Metadata;
}

export interface CoverUpdate {
  update: boolean;
  bytes: Uint8Array;
  mimeType: string;
}

export interface CoverReq {
  data: string;
  type: string;
}

export interface Metadata {
  title: string;
  artist: string;
  album: string;
  year: number;
  genre: string;
  // coverUrl: string;
}
