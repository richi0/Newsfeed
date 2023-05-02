export type Feed = {
  ID: number;
  Url: string;
  Title: string;
  Link: string;
  Description: string;
  Language: string;
  Copyright: string;
  ManagingEditor: string;
  WebMaster: string;
  PubDate: string;
  LastBuildDate: string;
  Category: string;
  Generator: string;
  Docs: string;
  CloudDomain: string;
  CloudPort: number;
  CloudPath: string;
  CloudRegisterProcedure: string;
  CloudProtocol: string;
  TTL: number;
  ImageUrl: string;
  ImageTitle: string;
  ImageLink: string;
  ImageWidth: number;
  ImageHeight: number;
  ImageDescription: string;
  Rating: string;
  TextInputTitle: string;
  TextInputDescription: string;
  TextInputName: string;
  TextInputLink: string;
  SkipHours: string;
  SkipDays: string;
  LastFetched: string;
  CreatedAt: string;
  UpdatedAt: string;
};

export type News = {
  ID: number;
  Title: string;
  Link: string;
  Description: string;
  Author: string;
  Category: string;
  Comments: string;
  EnclosureUrl: string;
  EnclosureLength: number;
  EnclosureType: string;
  GuidUrl: string;
  GuidIsPermaLink: string;
  SourceUrl: string;
  SourceText: string;
  PubDate: string;
  FeedID: number;
  CreatedAt: string;
  UpdatedAt: string;
};