export interface MusicFile {
  uid: string;
  name: string;
  size: number;
  metadata: Metadata;
}

export interface Metadata {
  title: string;
  artist: string;
  album: string;
  year: string;
  genre: string;
  // coverUrl: string;
}
