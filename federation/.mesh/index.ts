// @ts-nocheck
import { GraphQLResolveInfo, SelectionSetNode, FieldNode, GraphQLScalarType, GraphQLScalarTypeConfig } from 'graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
import { gql } from '@graphql-mesh/utils';

import { findAndParseConfig } from '@graphql-mesh/cli';
import { createMeshHTTPHandler, MeshHTTPHandler } from '@graphql-mesh/http';
import { getMesh, ExecuteMeshFn, SubscribeMeshFn, MeshContext as BaseMeshContext, MeshInstance } from '@graphql-mesh/runtime';
import { MeshStore, FsStoreStorageAdapter } from '@graphql-mesh/store';
import { path as pathModule } from '@graphql-mesh/cross-helpers';
import { ImportFn } from '@graphql-mesh/types';
import type { ExampleTypes } from './sources/Example/types';
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

export type WithIndex<TObject> = TObject & Record<string, any>;
export type ResolversObject<TObject> = WithIndex<TObject>;

export type ResolverTypeWrapper<T> = Promise<T> | T;


export type ResolverWithResolve<TResult, TParent, TContext, TArgs> = {
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};

export type LegacyStitchingResolver<TResult, TParent, TContext, TArgs> = {
  fragment: string;
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};

export type NewStitchingResolver<TResult, TParent, TContext, TArgs> = {
  selectionSet: string | ((fieldNode: FieldNode) => SelectionSetNode);
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};
export type StitchingResolver<TResult, TParent, TContext, TArgs> = LegacyStitchingResolver<TResult, TParent, TContext, TArgs> | NewStitchingResolver<TResult, TParent, TContext, TArgs>;
export type Resolver<TResult, TParent = {}, TContext = {}, TArgs = {}> =
  | ResolverFn<TResult, TParent, TContext, TArgs>
  | ResolverWithResolve<TResult, TParent, TContext, TArgs>
  | StitchingResolver<TResult, TParent, TContext, TArgs>;

export type ResolverFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => Promise<TResult> | TResult;

export type SubscriptionSubscribeFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => AsyncIterable<TResult> | Promise<AsyncIterable<TResult>>;

export type SubscriptionResolveFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;

export interface SubscriptionSubscriberObject<TResult, TKey extends string, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<{ [key in TKey]: TResult }, TParent, TContext, TArgs>;
  resolve?: SubscriptionResolveFn<TResult, { [key in TKey]: TResult }, TContext, TArgs>;
}

export interface SubscriptionResolverObject<TResult, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<any, TParent, TContext, TArgs>;
  resolve: SubscriptionResolveFn<TResult, any, TContext, TArgs>;
}

export type SubscriptionObject<TResult, TKey extends string, TParent, TContext, TArgs> =
  | SubscriptionSubscriberObject<TResult, TKey, TParent, TContext, TArgs>
  | SubscriptionResolverObject<TResult, TParent, TContext, TArgs>;

export type SubscriptionResolver<TResult, TKey extends string, TParent = {}, TContext = {}, TArgs = {}> =
  | ((...args: any[]) => SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>)
  | SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>;

export type TypeResolveFn<TTypes, TParent = {}, TContext = {}> = (
  parent: TParent,
  context: TContext,
  info: GraphQLResolveInfo
) => Maybe<TTypes> | Promise<Maybe<TTypes>>;

export type IsTypeOfResolverFn<T = {}, TContext = {}> = (obj: T, context: TContext, info: GraphQLResolveInfo) => boolean | Promise<boolean>;

export type NextResolverFn<T> = () => Promise<T>;

export type DirectiveResolverFn<TResult = {}, TParent = {}, TContext = {}, TArgs = {}> = (
  next: NextResolverFn<TResult>,
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;



/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = ResolversObject<{
  Query: ResolverTypeWrapper<{}>;
  MoviesResult: ResolverTypeWrapper<MoviesResult>;
  Movie: ResolverTypeWrapper<Movie>;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  BigInt: ResolverTypeWrapper<Scalars['BigInt']['output']>;
  Float: ResolverTypeWrapper<Scalars['Float']['output']>;
  google_protobuf_Timestamp: ResolverTypeWrapper<google_protobuf_Timestamp>;
  Int: ResolverTypeWrapper<Scalars['Int']['output']>;
  Genre: Genre;
  MovieRequest_Input: MovieRequest_Input;
  Movie_Input: Movie_Input;
  google_protobuf_Timestamp_Input: google_protobuf_Timestamp_Input;
  SearchByCastRequest_Input: SearchByCastRequest_Input;
  ConnectivityState: ConnectivityState;
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  ObjMap: ResolverTypeWrapper<Scalars['ObjMap']['output']>;
}>;

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = ResolversObject<{
  Query: {};
  MoviesResult: MoviesResult;
  Movie: Movie;
  String: Scalars['String']['output'];
  BigInt: Scalars['BigInt']['output'];
  Float: Scalars['Float']['output'];
  google_protobuf_Timestamp: google_protobuf_Timestamp;
  Int: Scalars['Int']['output'];
  MovieRequest_Input: MovieRequest_Input;
  Movie_Input: Movie_Input;
  google_protobuf_Timestamp_Input: google_protobuf_Timestamp_Input;
  SearchByCastRequest_Input: SearchByCastRequest_Input;
  Boolean: Scalars['Boolean']['output'];
  ObjMap: Scalars['ObjMap']['output'];
}>;

export type grpcMethodDirectiveArgs = {
  rootJsonName?: Maybe<Scalars['String']['input']>;
  objPath?: Maybe<Scalars['String']['input']>;
  methodName?: Maybe<Scalars['String']['input']>;
  responseStream?: Maybe<Scalars['Boolean']['input']>;
};

export type grpcMethodDirectiveResolver<Result, Parent, ContextType = MeshContext, Args = grpcMethodDirectiveArgs> = DirectiveResolverFn<Result, Parent, ContextType, Args>;

export type grpcConnectivityStateDirectiveArgs = {
  rootJsonName?: Maybe<Scalars['String']['input']>;
  objPath?: Maybe<Scalars['String']['input']>;
};

export type grpcConnectivityStateDirectiveResolver<Result, Parent, ContextType = MeshContext, Args = grpcConnectivityStateDirectiveArgs> = DirectiveResolverFn<Result, Parent, ContextType, Args>;

export type grpcRootJsonDirectiveArgs = {
  name?: Maybe<Scalars['String']['input']>;
  rootJson?: Maybe<Scalars['ObjMap']['input']>;
  loadOptions?: Maybe<Scalars['ObjMap']['input']>;
};

export type grpcRootJsonDirectiveResolver<Result, Parent, ContextType = MeshContext, Args = grpcRootJsonDirectiveArgs> = DirectiveResolverFn<Result, Parent, ContextType, Args>;

export type streamDirectiveArgs = {
  if?: Scalars['Boolean']['input'];
  label?: Maybe<Scalars['String']['input']>;
  initialCount?: Maybe<Scalars['Int']['input']>;
};

export type streamDirectiveResolver<Result, Parent, ContextType = MeshContext, Args = streamDirectiveArgs> = DirectiveResolverFn<Result, Parent, ContextType, Args>;

export type QueryResolvers<ContextType = MeshContext, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = ResolversObject<{
  exampleGetMovies?: Resolver<Maybe<ResolversTypes['MoviesResult']>, ParentType, ContextType, Partial<QueryexampleGetMoviesArgs>>;
  exampleSearchMoviesByCast?: Resolver<Maybe<Array<Maybe<ResolversTypes['Movie']>>>, ParentType, ContextType, Partial<QueryexampleSearchMoviesByCastArgs>>;
  exampleConnectivityState?: Resolver<Maybe<ResolversTypes['ConnectivityState']>, ParentType, ContextType, Partial<QueryexampleConnectivityStateArgs>>;
  anotherExampleGetMovies?: Resolver<Maybe<ResolversTypes['MoviesResult']>, ParentType, ContextType, Partial<QueryanotherExampleGetMoviesArgs>>;
  anotherExampleSearchMoviesByCast?: Resolver<Maybe<Array<Maybe<ResolversTypes['Movie']>>>, ParentType, ContextType, Partial<QueryanotherExampleSearchMoviesByCastArgs>>;
  anotherExampleConnectivityState?: Resolver<Maybe<ResolversTypes['ConnectivityState']>, ParentType, ContextType, Partial<QueryanotherExampleConnectivityStateArgs>>;
}>;

export type MoviesResultResolvers<ContextType = MeshContext, ParentType extends ResolversParentTypes['MoviesResult'] = ResolversParentTypes['MoviesResult']> = ResolversObject<{
  result?: Resolver<Maybe<Array<Maybe<ResolversTypes['Movie']>>>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type MovieResolvers<ContextType = MeshContext, ParentType extends ResolversParentTypes['Movie'] = ResolversParentTypes['Movie']> = ResolversObject<{
  name?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  year?: Resolver<Maybe<ResolversTypes['BigInt']>, ParentType, ContextType>;
  rating?: Resolver<Maybe<ResolversTypes['Float']>, ParentType, ContextType>;
  cast?: Resolver<Maybe<Array<Maybe<ResolversTypes['String']>>>, ParentType, ContextType>;
  time?: Resolver<Maybe<ResolversTypes['google_protobuf_Timestamp']>, ParentType, ContextType>;
  genre?: Resolver<Maybe<ResolversTypes['Genre']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export interface BigIntScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['BigInt'], any> {
  name: 'BigInt';
}

export type google_protobuf_TimestampResolvers<ContextType = MeshContext, ParentType extends ResolversParentTypes['google_protobuf_Timestamp'] = ResolversParentTypes['google_protobuf_Timestamp']> = ResolversObject<{
  seconds?: Resolver<Maybe<ResolversTypes['BigInt']>, ParentType, ContextType>;
  nanos?: Resolver<Maybe<ResolversTypes['Int']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
}>;

export type GenreResolvers = { UNSPECIFIED: 0, ACTION: 1, DRAMA: 2 };

export type ConnectivityStateResolvers = { IDLE: 0, CONNECTING: 1, READY: 2, TRANSIENT_FAILURE: 3, SHUTDOWN: 4 };

export interface ObjMapScalarConfig extends GraphQLScalarTypeConfig<ResolversTypes['ObjMap'], any> {
  name: 'ObjMap';
}

export type Resolvers<ContextType = MeshContext> = ResolversObject<{
  Query?: QueryResolvers<ContextType>;
  MoviesResult?: MoviesResultResolvers<ContextType>;
  Movie?: MovieResolvers<ContextType>;
  BigInt?: GraphQLScalarType;
  google_protobuf_Timestamp?: google_protobuf_TimestampResolvers<ContextType>;
  Genre?: GenreResolvers;
  ConnectivityState?: ConnectivityStateResolvers;
  ObjMap?: GraphQLScalarType;
}>;

export type DirectiveResolvers<ContextType = MeshContext> = ResolversObject<{
  grpcMethod?: grpcMethodDirectiveResolver<any, any, ContextType>;
  grpcConnectivityState?: grpcConnectivityStateDirectiveResolver<any, any, ContextType>;
  grpcRootJson?: grpcRootJsonDirectiveResolver<any, any, ContextType>;
  stream?: streamDirectiveResolver<any, any, ContextType>;
}>;

export type MeshContext = ExampleTypes.Context & BaseMeshContext;


const baseDir = pathModule.join(typeof __dirname === 'string' ? __dirname : '/', '..');

const importFn: ImportFn = <T>(moduleId: string) => {
  const relativeModuleId = (pathModule.isAbsolute(moduleId) ? pathModule.relative(baseDir, moduleId) : moduleId).split('\\').join('/').replace(baseDir + '/', '');
  switch(relativeModuleId) {
    default:
      return Promise.reject(new Error(`Cannot find module '${relativeModuleId}'.`));
  }
};

const rootStore = new MeshStore('.mesh', new FsStoreStorageAdapter({
  cwd: baseDir,
  importFn,
  fileType: "ts",
}), {
  readonly: true,
  validate: false
});

export function getMeshOptions() {
  console.warn('WARNING: These artifacts are built for development mode. Please run "mesh build" to build production artifacts');
  return findAndParseConfig({
    dir: baseDir,
    artifactsDir: ".mesh",
    configName: "mesh",
    additionalPackagePrefixes: [],
    initialLoggerPrefix: "🕸️  Mesh",
  });
}

export function createBuiltMeshHTTPHandler<TServerContext = {}>(): MeshHTTPHandler<TServerContext> {
  return createMeshHTTPHandler<TServerContext>({
    baseDir,
    getBuiltMesh: getBuiltMesh,
    rawServeConfig: undefined,
  })
}

let meshInstance$: Promise<MeshInstance> | undefined;

export function getBuiltMesh(): Promise<MeshInstance> {
  if (meshInstance$ == null) {
    meshInstance$ = getMeshOptions().then(meshOptions => getMesh(meshOptions)).then(mesh => {
      const id = mesh.pubsub.subscribe('destroy', () => {
        meshInstance$ = undefined;
        mesh.pubsub.unsubscribe(id);
      });
      return mesh;
    });
  }
  return meshInstance$;
}

export const execute: ExecuteMeshFn = (...args) => getBuiltMesh().then(({ execute }) => execute(...args));

export const subscribe: SubscribeMeshFn = (...args) => getBuiltMesh().then(({ subscribe }) => subscribe(...args));
export function getMeshSDK<TGlobalContext = any, TOperationContext = any>(globalContext?: TGlobalContext) {
  const sdkRequester$ = getBuiltMesh().then(({ sdkRequesterFactory }) => sdkRequesterFactory(globalContext));
  return getSdk<TOperationContext, TGlobalContext>((...args) => sdkRequester$.then(sdkRequester => sdkRequester(...args)));
}
export type GetMoviesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetMoviesQuery = { exampleGetMovies?: Maybe<{ result?: Maybe<Array<Maybe<(
      Pick<Movie, 'name' | 'year' | 'rating' | 'cast'>
      & { time?: Maybe<Pick<google_protobuf_Timestamp, 'seconds'>> }
    )>>> }> };

export type SearchMoviesByCastQueryVariables = Exact<{ [key: string]: never; }>;


export type SearchMoviesByCastQuery = { exampleSearchMoviesByCast?: Maybe<Array<Maybe<Pick<Movie, 'name' | 'year' | 'rating' | 'cast'>>>> };


export const GetMoviesDocument = gql`
    query GetMovies {
  exampleGetMovies(input: {movie: {genre: DRAMA, year: 2015}}) {
    result {
      name
      year
      rating
      cast
      time {
        seconds
      }
    }
  }
}
    ` as unknown as DocumentNode<GetMoviesQuery, GetMoviesQueryVariables>;
export const SearchMoviesByCastDocument = gql`
    query SearchMoviesByCast {
  exampleSearchMoviesByCast(input: {castName: "Tom Cruise"}) @stream {
    name
    year
    rating
    cast
  }
}
    ` as unknown as DocumentNode<SearchMoviesByCastQuery, SearchMoviesByCastQueryVariables>;



export type Requester<C = {}, E = unknown> = <R, V>(doc: DocumentNode, vars?: V, options?: C) => Promise<R> | AsyncIterable<R>
export function getSdk<C, E>(requester: Requester<C, E>) {
  return {
    GetMovies(variables?: GetMoviesQueryVariables, options?: C): Promise<GetMoviesQuery> {
      return requester<GetMoviesQuery, GetMoviesQueryVariables>(GetMoviesDocument, variables, options) as Promise<GetMoviesQuery>;
    },
    SearchMoviesByCast(variables?: SearchMoviesByCastQueryVariables, options?: C): Promise<SearchMoviesByCastQuery> {
      return requester<SearchMoviesByCastQuery, SearchMoviesByCastQueryVariables>(SearchMoviesByCastDocument, variables, options) as Promise<SearchMoviesByCastQuery>;
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;