// @ts-nocheck

import { InContextSdkMethod } from '@graphql-mesh/types';
import { MeshContext } from '@graphql-mesh/runtime';

export namespace ExampleTypes {
  export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  /** The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text. */
  String: { input: string; output: string; }
  /** The `Boolean` scalar type represents `true` or `false`. */
  Boolean: { input: boolean; output: boolean; }
  /** The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1. */
  Int: { input: number; output: number; }
  /** The `Float` scalar type represents signed double-precision fractional values as specified by [IEEE 754](https://en.wikipedia.org/wiki/IEEE_floating_point). */
  Float: { input: number; output: number; }
  /** The `BigInt` scalar type represents non-fractional signed whole numeric values. */
  BigInt: { input: bigint; output: bigint; }
  ObjMap: { input: any; output: any; }
};

export type Query = {
  /** get all movies */
  exampleGetMovies?: Maybe<MoviesResult>;
  /** search movies by the name of the cast */
  exampleSearchMoviesByCast?: Maybe<Array<Maybe<Movie>>>;
  exampleConnectivityState?: Maybe<ConnectivityState>;
  /** get all movies */
  anotherExampleGetMovies?: Maybe<MoviesResult>;
  /** search movies by the name of the cast */
  anotherExampleSearchMoviesByCast?: Maybe<Array<Maybe<Movie>>>;
  anotherExampleConnectivityState?: Maybe<ConnectivityState>;
};


export type QueryexampleGetMoviesArgs = {
  input?: InputMaybe<MovieRequest_Input>;
};


export type QueryexampleSearchMoviesByCastArgs = {
  input?: InputMaybe<SearchByCastRequest_Input>;
};


export type QueryexampleConnectivityStateArgs = {
  tryToConnect?: InputMaybe<Scalars['Boolean']['input']>;
};


export type QueryanotherExampleGetMoviesArgs = {
  input?: InputMaybe<MovieRequest_Input>;
};


export type QueryanotherExampleSearchMoviesByCastArgs = {
  input?: InputMaybe<SearchByCastRequest_Input>;
};


export type QueryanotherExampleConnectivityStateArgs = {
  tryToConnect?: InputMaybe<Scalars['Boolean']['input']>;
};

/** movie result message, contains list of movies */
export type MoviesResult = {
  /** list of movies */
  result?: Maybe<Array<Maybe<Movie>>>;
};

/** movie message payload */
export type Movie = {
  name?: Maybe<Scalars['String']['output']>;
  year?: Maybe<Scalars['BigInt']['output']>;
  rating?: Maybe<Scalars['Float']['output']>;
  /** list of cast */
  cast?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  time?: Maybe<google_protobuf_Timestamp>;
  genre?: Maybe<Genre>;
};

export type google_protobuf_Timestamp = {
  seconds?: Maybe<Scalars['BigInt']['output']>;
  nanos?: Maybe<Scalars['Int']['output']>;
};

export type Genre =
  | 'UNSPECIFIED'
  | 'ACTION'
  | 'DRAMA';

export type MovieRequest_Input = {
  movie?: InputMaybe<Movie_Input>;
};

/** movie message payload */
export type Movie_Input = {
  name?: InputMaybe<Scalars['String']['input']>;
  year?: InputMaybe<Scalars['BigInt']['input']>;
  rating?: InputMaybe<Scalars['Float']['input']>;
  /** list of cast */
  cast?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  time?: InputMaybe<google_protobuf_Timestamp_Input>;
  genre?: InputMaybe<Genre>;
};

export type google_protobuf_Timestamp_Input = {
  seconds?: InputMaybe<Scalars['BigInt']['input']>;
  nanos?: InputMaybe<Scalars['Int']['input']>;
};

export type SearchByCastRequest_Input = {
  castName?: InputMaybe<Scalars['String']['input']>;
};

export type ConnectivityState =
  | 'IDLE'
  | 'CONNECTING'
  | 'READY'
  | 'TRANSIENT_FAILURE'
  | 'SHUTDOWN';

  export type QuerySdk = {
      /** get all movies **/
  exampleGetMovies: InContextSdkMethod<Query['exampleGetMovies'], QueryexampleGetMoviesArgs, MeshContext>,
  /** search movies by the name of the cast **/
  exampleSearchMoviesByCast: InContextSdkMethod<Query['exampleSearchMoviesByCast'], QueryexampleSearchMoviesByCastArgs, MeshContext>,
  /** undefined **/
  exampleConnectivityState: InContextSdkMethod<Query['exampleConnectivityState'], QueryexampleConnectivityStateArgs, MeshContext>,
  /** get all movies **/
  anotherExampleGetMovies: InContextSdkMethod<Query['anotherExampleGetMovies'], QueryanotherExampleGetMoviesArgs, MeshContext>,
  /** search movies by the name of the cast **/
  anotherExampleSearchMoviesByCast: InContextSdkMethod<Query['anotherExampleSearchMoviesByCast'], QueryanotherExampleSearchMoviesByCastArgs, MeshContext>,
  /** undefined **/
  anotherExampleConnectivityState: InContextSdkMethod<Query['anotherExampleConnectivityState'], QueryanotherExampleConnectivityStateArgs, MeshContext>
  };

  export type MutationSdk = {
    
  };

  export type SubscriptionSdk = {
    
  };

  export type Context = {
      ["Example"]: { Query: QuerySdk, Mutation: MutationSdk, Subscription: SubscriptionSdk },
      
    };
}
