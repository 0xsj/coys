import protoloader from '@grpc/proto-loader';
import { credentials, loadPackageDefinition } from '@grpc/grpc-js';

const accountPackageDefinition = protoloader.loadSync('./src/lib/proto/account.proto', {});

const accountProto = loadPackageDefinition(accountPackageDefinition);

const accountClient = new accountProto.proto.AccountService();
