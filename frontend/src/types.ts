export interface MusicFile {
  uid: string;
  name: string;
  size: number;
  metadata: Metadata;
}

export interface CoverUpdate {
  update: boolean;
  bytes: Uint8Array;
}

export interface Metadata {
  title: string;
  artist: string;
  album: string;
  year: string;
  genre: string;
  // coverUrl: string;
}
