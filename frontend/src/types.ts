export interface MusicFile {
  id: string;
  filename: string;
  format: string;
  size: number;
  path: string;
  artist?: string;
  album?: string;
  title?: string;
  year?: string;
  genre?: string;
  trackNumber?: string;
  bitrate?: number;
  duration: number;
  sampleRate?: number;
}
