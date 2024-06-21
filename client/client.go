// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateDocumentAnalyzeTaskRequest struct {
	Document *CreateDocumentAnalyzeTaskRequestDocument `json:"document,omitempty" xml:"document,omitempty" type:"Struct"`
	Output   *CreateDocumentAnalyzeTaskRequestOutput   `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
}

func (s CreateDocumentAnalyzeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskRequest) SetDocument(v *CreateDocumentAnalyzeTaskRequestDocument) *CreateDocumentAnalyzeTaskRequest {
	s.Document = v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequest) SetOutput(v *CreateDocumentAnalyzeTaskRequestOutput) *CreateDocumentAnalyzeTaskRequest {
	s.Output = v
	return s
}

type CreateDocumentAnalyzeTaskRequestDocument struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	FileType *string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateDocumentAnalyzeTaskRequestDocument) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskRequestDocument) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetContent(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.Content = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetFileName(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.FileName = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetFileType(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.FileType = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetUrl(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.Url = &v
	return s
}

type CreateDocumentAnalyzeTaskRequestOutput struct {
	ImageStorage *string `json:"image_storage,omitempty" xml:"image_storage,omitempty"`
}

func (s CreateDocumentAnalyzeTaskRequestOutput) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskRequestOutput) SetImageStorage(v string) *CreateDocumentAnalyzeTaskRequestOutput {
	s.ImageStorage = &v
	return s
}

type CreateDocumentAnalyzeTaskResponseBody struct {
	Latency   *int32                                       `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                      `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateDocumentAnalyzeTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateDocumentAnalyzeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskResponseBody) SetLatency(v int32) *CreateDocumentAnalyzeTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponseBody) SetRequestId(v string) *CreateDocumentAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponseBody) SetResult(v *CreateDocumentAnalyzeTaskResponseBodyResult) *CreateDocumentAnalyzeTaskResponseBody {
	s.Result = v
	return s
}

type CreateDocumentAnalyzeTaskResponseBodyResult struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateDocumentAnalyzeTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskResponseBodyResult) SetTaskId(v string) *CreateDocumentAnalyzeTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

type CreateDocumentAnalyzeTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocumentAnalyzeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocumentAnalyzeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskResponse) SetHeaders(v map[string]*string) *CreateDocumentAnalyzeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponse) SetStatusCode(v int32) *CreateDocumentAnalyzeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponse) SetBody(v *CreateDocumentAnalyzeTaskResponseBody) *CreateDocumentAnalyzeTaskResponse {
	s.Body = v
	return s
}

type GetDocumentAnalyzeTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusRequest) SetTaskId(v string) *GetDocumentAnalyzeTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetDocumentAnalyzeTaskStatusResponseBody struct {
	Latency   *int32                                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetDocumentAnalyzeTaskStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetDocumentAnalyzeTaskStatusResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetLatency(v int32) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.Latency = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetRequestId(v string) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetResult(v *GetDocumentAnalyzeTaskStatusResponseBodyResult) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBody) SetUsage(v *GetDocumentAnalyzeTaskStatusResponseBodyUsage) *GetDocumentAnalyzeTaskStatusResponseBody {
	s.Usage = v
	return s
}

type GetDocumentAnalyzeTaskStatusResponseBodyResult struct {
	Data   *GetDocumentAnalyzeTaskStatusResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Status *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string                                             `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) SetData(v *GetDocumentAnalyzeTaskStatusResponseBodyResultData) *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) SetStatus(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResult) SetTaskId(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

type GetDocumentAnalyzeTaskStatusResponseBodyResultData struct {
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	PageNum     *int32  `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) SetContent(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResultData {
	s.Content = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) SetContentType(v string) *GetDocumentAnalyzeTaskStatusResponseBodyResultData {
	s.ContentType = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyResultData) SetPageNum(v int32) *GetDocumentAnalyzeTaskStatusResponseBodyResultData {
	s.PageNum = &v
	return s
}

type GetDocumentAnalyzeTaskStatusResponseBodyUsage struct {
	ImageCount *int64 `json:"image_count,omitempty" xml:"image_count,omitempty"`
	TableCount *int64 `json:"table_count,omitempty" xml:"table_count,omitempty"`
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) SetImageCount(v int64) *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	s.ImageCount = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) SetTableCount(v int64) *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	s.TableCount = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponseBodyUsage) SetTokenCount(v int64) *GetDocumentAnalyzeTaskStatusResponseBodyUsage {
	s.TokenCount = &v
	return s
}

type GetDocumentAnalyzeTaskStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentAnalyzeTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentAnalyzeTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentAnalyzeTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentAnalyzeTaskStatusResponse) SetHeaders(v map[string]*string) *GetDocumentAnalyzeTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponse) SetStatusCode(v int32) *GetDocumentAnalyzeTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentAnalyzeTaskStatusResponse) SetBody(v *GetDocumentAnalyzeTaskStatusResponseBody) *GetDocumentAnalyzeTaskStatusResponse {
	s.Body = v
	return s
}

type GetDocumentRankRequest struct {
	// This parameter is required.
	Docs []*string `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetDocumentRankRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentRankRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentRankRequest) SetDocs(v []*string) *GetDocumentRankRequest {
	s.Docs = v
	return s
}

func (s *GetDocumentRankRequest) SetQuery(v string) *GetDocumentRankRequest {
	s.Query = &v
	return s
}

type GetDocumentRankResponseBody struct {
	Latency   *int32                             `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                            `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetDocumentRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetDocumentRankResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetDocumentRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentRankResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBody) SetLatency(v int32) *GetDocumentRankResponseBody {
	s.Latency = &v
	return s
}

func (s *GetDocumentRankResponseBody) SetRequestId(v string) *GetDocumentRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentRankResponseBody) SetResult(v *GetDocumentRankResponseBodyResult) *GetDocumentRankResponseBody {
	s.Result = v
	return s
}

func (s *GetDocumentRankResponseBody) SetUsage(v *GetDocumentRankResponseBodyUsage) *GetDocumentRankResponseBody {
	s.Usage = v
	return s
}

type GetDocumentRankResponseBodyResult struct {
	Scores []*GetDocumentRankResponseBodyResultScores `json:"scores,omitempty" xml:"scores,omitempty" type:"Repeated"`
}

func (s GetDocumentRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBodyResult) SetScores(v []*GetDocumentRankResponseBodyResultScores) *GetDocumentRankResponseBodyResult {
	s.Scores = v
	return s
}

type GetDocumentRankResponseBodyResultScores struct {
	Index *int32   `json:"index,omitempty" xml:"index,omitempty"`
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
}

func (s GetDocumentRankResponseBodyResultScores) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentRankResponseBodyResultScores) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBodyResultScores) SetIndex(v int32) *GetDocumentRankResponseBodyResultScores {
	s.Index = &v
	return s
}

func (s *GetDocumentRankResponseBodyResultScores) SetScore(v float64) *GetDocumentRankResponseBodyResultScores {
	s.Score = &v
	return s
}

type GetDocumentRankResponseBodyUsage struct {
	DocCount *int64 `json:"doc_count,omitempty" xml:"doc_count,omitempty"`
}

func (s GetDocumentRankResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentRankResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBodyUsage) SetDocCount(v int64) *GetDocumentRankResponseBodyUsage {
	s.DocCount = &v
	return s
}

type GetDocumentRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentRankResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentRankResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponse) SetHeaders(v map[string]*string) *GetDocumentRankResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentRankResponse) SetStatusCode(v int32) *GetDocumentRankResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentRankResponse) SetBody(v *GetDocumentRankResponseBody) *GetDocumentRankResponse {
	s.Body = v
	return s
}

type GetDocumentSplitRequest struct {
	// This parameter is required.
	Document *GetDocumentSplitRequestDocument `json:"document,omitempty" xml:"document,omitempty" type:"Struct"`
	Strategy *GetDocumentSplitRequestStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
}

func (s GetDocumentSplitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitRequest) SetDocument(v *GetDocumentSplitRequestDocument) *GetDocumentSplitRequest {
	s.Document = v
	return s
}

func (s *GetDocumentSplitRequest) SetStrategy(v *GetDocumentSplitRequestStrategy) *GetDocumentSplitRequest {
	s.Strategy = v
	return s
}

type GetDocumentSplitRequestDocument struct {
	Content         *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentEncoding *string `json:"content_encoding,omitempty" xml:"content_encoding,omitempty"`
	ContentType     *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
}

func (s GetDocumentSplitRequestDocument) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitRequestDocument) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitRequestDocument) SetContent(v string) *GetDocumentSplitRequestDocument {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitRequestDocument) SetContentEncoding(v string) *GetDocumentSplitRequestDocument {
	s.ContentEncoding = &v
	return s
}

func (s *GetDocumentSplitRequestDocument) SetContentType(v string) *GetDocumentSplitRequestDocument {
	s.ContentType = &v
	return s
}

type GetDocumentSplitRequestStrategy struct {
	ComputeType  *string `json:"compute_type,omitempty" xml:"compute_type,omitempty"`
	MaxChunkSize *int64  `json:"max_chunk_size,omitempty" xml:"max_chunk_size,omitempty"`
	NeedSentence *bool   `json:"need_sentence,omitempty" xml:"need_sentence,omitempty"`
}

func (s GetDocumentSplitRequestStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitRequestStrategy) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitRequestStrategy) SetComputeType(v string) *GetDocumentSplitRequestStrategy {
	s.ComputeType = &v
	return s
}

func (s *GetDocumentSplitRequestStrategy) SetMaxChunkSize(v int64) *GetDocumentSplitRequestStrategy {
	s.MaxChunkSize = &v
	return s
}

func (s *GetDocumentSplitRequestStrategy) SetNeedSentence(v bool) *GetDocumentSplitRequestStrategy {
	s.NeedSentence = &v
	return s
}

type GetDocumentSplitResponseBody struct {
	Latency   *int32                              `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                             `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetDocumentSplitResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetDocumentSplitResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetDocumentSplitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBody) SetLatency(v int32) *GetDocumentSplitResponseBody {
	s.Latency = &v
	return s
}

func (s *GetDocumentSplitResponseBody) SetRequestId(v string) *GetDocumentSplitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentSplitResponseBody) SetResult(v *GetDocumentSplitResponseBodyResult) *GetDocumentSplitResponseBody {
	s.Result = v
	return s
}

func (s *GetDocumentSplitResponseBody) SetUsage(v *GetDocumentSplitResponseBodyUsage) *GetDocumentSplitResponseBody {
	s.Usage = v
	return s
}

type GetDocumentSplitResponseBodyResult struct {
	Chunks    []*GetDocumentSplitResponseBodyResultChunks    `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	Nodes     []map[string]*string                           `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	RichTexts []*GetDocumentSplitResponseBodyResultRichTexts `json:"rich_texts,omitempty" xml:"rich_texts,omitempty" type:"Repeated"`
	Sentences []*GetDocumentSplitResponseBodyResultSentences `json:"sentences,omitempty" xml:"sentences,omitempty" type:"Repeated"`
}

func (s GetDocumentSplitResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResult) SetChunks(v []*GetDocumentSplitResponseBodyResultChunks) *GetDocumentSplitResponseBodyResult {
	s.Chunks = v
	return s
}

func (s *GetDocumentSplitResponseBodyResult) SetNodes(v []map[string]*string) *GetDocumentSplitResponseBodyResult {
	s.Nodes = v
	return s
}

func (s *GetDocumentSplitResponseBodyResult) SetRichTexts(v []*GetDocumentSplitResponseBodyResultRichTexts) *GetDocumentSplitResponseBodyResult {
	s.RichTexts = v
	return s
}

func (s *GetDocumentSplitResponseBodyResult) SetSentences(v []*GetDocumentSplitResponseBodyResultSentences) *GetDocumentSplitResponseBodyResult {
	s.Sentences = v
	return s
}

type GetDocumentSplitResponseBodyResultChunks struct {
	Content *string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta    map[string]*string `json:"meta,omitempty" xml:"meta,omitempty"`
}

func (s GetDocumentSplitResponseBodyResultChunks) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResultChunks) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResultChunks) SetContent(v string) *GetDocumentSplitResponseBodyResultChunks {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitResponseBodyResultChunks) SetMeta(v map[string]*string) *GetDocumentSplitResponseBodyResultChunks {
	s.Meta = v
	return s
}

type GetDocumentSplitResponseBodyResultRichTexts struct {
	Content *string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta    map[string]*string `json:"meta,omitempty" xml:"meta,omitempty"`
}

func (s GetDocumentSplitResponseBodyResultRichTexts) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResultRichTexts) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResultRichTexts) SetContent(v string) *GetDocumentSplitResponseBodyResultRichTexts {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitResponseBodyResultRichTexts) SetMeta(v map[string]*string) *GetDocumentSplitResponseBodyResultRichTexts {
	s.Meta = v
	return s
}

type GetDocumentSplitResponseBodyResultSentences struct {
	Content *string            `json:"content,omitempty" xml:"content,omitempty"`
	Meta    map[string]*string `json:"meta,omitempty" xml:"meta,omitempty"`
}

func (s GetDocumentSplitResponseBodyResultSentences) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitResponseBodyResultSentences) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyResultSentences) SetContent(v string) *GetDocumentSplitResponseBodyResultSentences {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitResponseBodyResultSentences) SetMeta(v map[string]*string) *GetDocumentSplitResponseBodyResultSentences {
	s.Meta = v
	return s
}

type GetDocumentSplitResponseBodyUsage struct {
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetDocumentSplitResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponseBodyUsage) SetTokenCount(v int64) *GetDocumentSplitResponseBodyUsage {
	s.TokenCount = &v
	return s
}

type GetDocumentSplitResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentSplitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentSplitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSplitResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponse) SetHeaders(v map[string]*string) *GetDocumentSplitResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentSplitResponse) SetStatusCode(v int32) *GetDocumentSplitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentSplitResponse) SetBody(v *GetDocumentSplitResponseBody) *GetDocumentSplitResponse {
	s.Body = v
	return s
}

type GetTextEmbeddingRequest struct {
	// This parameter is required.
	Input []*string `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	// example:
	//
	// document
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s GetTextEmbeddingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTextEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingRequest) SetInput(v []*string) *GetTextEmbeddingRequest {
	s.Input = v
	return s
}

func (s *GetTextEmbeddingRequest) SetInputType(v string) *GetTextEmbeddingRequest {
	s.InputType = &v
	return s
}

type GetTextEmbeddingResponseBody struct {
	Latency   *int32                              `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                             `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetTextEmbeddingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetTextEmbeddingResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTextEmbeddingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTextEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBody) SetLatency(v int32) *GetTextEmbeddingResponseBody {
	s.Latency = &v
	return s
}

func (s *GetTextEmbeddingResponseBody) SetRequestId(v string) *GetTextEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextEmbeddingResponseBody) SetResult(v *GetTextEmbeddingResponseBodyResult) *GetTextEmbeddingResponseBody {
	s.Result = v
	return s
}

func (s *GetTextEmbeddingResponseBody) SetUsage(v *GetTextEmbeddingResponseBodyUsage) *GetTextEmbeddingResponseBody {
	s.Usage = v
	return s
}

type GetTextEmbeddingResponseBodyResult struct {
	Embeddings []*GetTextEmbeddingResponseBodyResultEmbeddings `json:"embeddings,omitempty" xml:"embeddings,omitempty" type:"Repeated"`
}

func (s GetTextEmbeddingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTextEmbeddingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBodyResult) SetEmbeddings(v []*GetTextEmbeddingResponseBodyResultEmbeddings) *GetTextEmbeddingResponseBodyResult {
	s.Embeddings = v
	return s
}

type GetTextEmbeddingResponseBodyResultEmbeddings struct {
	Embedding []*float64 `json:"embedding,omitempty" xml:"embedding,omitempty" type:"Repeated"`
	Index     *int32     `json:"index,omitempty" xml:"index,omitempty"`
}

func (s GetTextEmbeddingResponseBodyResultEmbeddings) String() string {
	return tea.Prettify(s)
}

func (s GetTextEmbeddingResponseBodyResultEmbeddings) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBodyResultEmbeddings) SetEmbedding(v []*float64) *GetTextEmbeddingResponseBodyResultEmbeddings {
	s.Embedding = v
	return s
}

func (s *GetTextEmbeddingResponseBodyResultEmbeddings) SetIndex(v int32) *GetTextEmbeddingResponseBodyResultEmbeddings {
	s.Index = &v
	return s
}

type GetTextEmbeddingResponseBodyUsage struct {
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetTextEmbeddingResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s GetTextEmbeddingResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponseBodyUsage) SetTokenCount(v int64) *GetTextEmbeddingResponseBodyUsage {
	s.TokenCount = &v
	return s
}

type GetTextEmbeddingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextEmbeddingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponse) SetHeaders(v map[string]*string) *GetTextEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *GetTextEmbeddingResponse) SetStatusCode(v int32) *GetTextEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextEmbeddingResponse) SetBody(v *GetTextEmbeddingResponseBody) *GetTextEmbeddingResponse {
	s.Body = v
	return s
}

type GetTextGenerationRequest struct {
	CsiLevel *string `json:"csi_level,omitempty" xml:"csi_level,omitempty"`
	// This parameter is required.
	Messages   []*GetTextGenerationRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	Parameters map[string]*string                  `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Stream     *bool                               `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s GetTextGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTextGenerationRequest) GoString() string {
	return s.String()
}

func (s *GetTextGenerationRequest) SetCsiLevel(v string) *GetTextGenerationRequest {
	s.CsiLevel = &v
	return s
}

func (s *GetTextGenerationRequest) SetMessages(v []*GetTextGenerationRequestMessages) *GetTextGenerationRequest {
	s.Messages = v
	return s
}

func (s *GetTextGenerationRequest) SetParameters(v map[string]*string) *GetTextGenerationRequest {
	s.Parameters = v
	return s
}

func (s *GetTextGenerationRequest) SetStream(v bool) *GetTextGenerationRequest {
	s.Stream = &v
	return s
}

type GetTextGenerationRequestMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetTextGenerationRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s GetTextGenerationRequestMessages) GoString() string {
	return s.String()
}

func (s *GetTextGenerationRequestMessages) SetContent(v string) *GetTextGenerationRequestMessages {
	s.Content = &v
	return s
}

func (s *GetTextGenerationRequestMessages) SetRole(v string) *GetTextGenerationRequestMessages {
	s.Role = &v
	return s
}

type GetTextGenerationResponseBody struct {
	Latency   *int32                               `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                              `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetTextGenerationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetTextGenerationResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTextGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTextGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponseBody) SetLatency(v int32) *GetTextGenerationResponseBody {
	s.Latency = &v
	return s
}

func (s *GetTextGenerationResponseBody) SetRequestId(v string) *GetTextGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextGenerationResponseBody) SetResult(v *GetTextGenerationResponseBodyResult) *GetTextGenerationResponseBody {
	s.Result = v
	return s
}

func (s *GetTextGenerationResponseBody) SetUsage(v *GetTextGenerationResponseBodyUsage) *GetTextGenerationResponseBody {
	s.Usage = v
	return s
}

type GetTextGenerationResponseBodyResult struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetTextGenerationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTextGenerationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponseBodyResult) SetText(v string) *GetTextGenerationResponseBodyResult {
	s.Text = &v
	return s
}

type GetTextGenerationResponseBodyUsage struct {
	InputTokens  *int64 `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
	OutputTokens *int64 `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
	TotalTokens  *int64 `json:"total_tokens,omitempty" xml:"total_tokens,omitempty"`
}

func (s GetTextGenerationResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s GetTextGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponseBodyUsage) SetInputTokens(v int64) *GetTextGenerationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *GetTextGenerationResponseBodyUsage) SetOutputTokens(v int64) *GetTextGenerationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetTextGenerationResponseBodyUsage) SetTotalTokens(v int64) *GetTextGenerationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

type GetTextGenerationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextGenerationResponse) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponse) SetHeaders(v map[string]*string) *GetTextGenerationResponse {
	s.Headers = v
	return s
}

func (s *GetTextGenerationResponse) SetStatusCode(v int32) *GetTextGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextGenerationResponse) SetBody(v *GetTextGenerationResponseBody) *GetTextGenerationResponse {
	s.Body = v
	return s
}

type GetTextSparseEmbeddingRequest struct {
	// This parameter is required.
	Input []*string `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	// example:
	//
	// document
	InputType   *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
	ReturnToken *bool   `json:"return_token,omitempty" xml:"return_token,omitempty"`
}

func (s GetTextSparseEmbeddingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTextSparseEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingRequest) SetInput(v []*string) *GetTextSparseEmbeddingRequest {
	s.Input = v
	return s
}

func (s *GetTextSparseEmbeddingRequest) SetInputType(v string) *GetTextSparseEmbeddingRequest {
	s.InputType = &v
	return s
}

func (s *GetTextSparseEmbeddingRequest) SetReturnToken(v bool) *GetTextSparseEmbeddingRequest {
	s.ReturnToken = &v
	return s
}

type GetTextSparseEmbeddingResponseBody struct {
	Latency   *int32                                    `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                   `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetTextSparseEmbeddingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetTextSparseEmbeddingResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTextSparseEmbeddingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBody) SetLatency(v int32) *GetTextSparseEmbeddingResponseBody {
	s.Latency = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBody) SetRequestId(v string) *GetTextSparseEmbeddingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBody) SetResult(v *GetTextSparseEmbeddingResponseBodyResult) *GetTextSparseEmbeddingResponseBody {
	s.Result = v
	return s
}

func (s *GetTextSparseEmbeddingResponseBody) SetUsage(v *GetTextSparseEmbeddingResponseBodyUsage) *GetTextSparseEmbeddingResponseBody {
	s.Usage = v
	return s
}

type GetTextSparseEmbeddingResponseBodyResult struct {
	SparseEmbeddings []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings `json:"sparse_embeddings,omitempty" xml:"sparse_embeddings,omitempty" type:"Repeated"`
}

func (s GetTextSparseEmbeddingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyResult) SetSparseEmbeddings(v []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) *GetTextSparseEmbeddingResponseBodyResult {
	s.SparseEmbeddings = v
	return s
}

type GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings struct {
	Embedding []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding `json:"embedding,omitempty" xml:"embedding,omitempty" type:"Repeated"`
	Index     *int32                                                               `json:"index,omitempty" xml:"index,omitempty"`
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) String() string {
	return tea.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) SetEmbedding(v []*GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings {
	s.Embedding = v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings) SetIndex(v int32) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddings {
	s.Index = &v
	return s
}

type GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding struct {
	Token   *string  `json:"token,omitempty" xml:"token,omitempty"`
	TokenId *int32   `json:"token_id,omitempty" xml:"token_id,omitempty"`
	Weight  *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) String() string {
	return tea.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) SetToken(v string) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding {
	s.Token = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) SetTokenId(v int32) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding {
	s.TokenId = &v
	return s
}

func (s *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding) SetWeight(v float32) *GetTextSparseEmbeddingResponseBodyResultSparseEmbeddingsEmbedding {
	s.Weight = &v
	return s
}

type GetTextSparseEmbeddingResponseBodyUsage struct {
	TokenCount *int64 `json:"token_count,omitempty" xml:"token_count,omitempty"`
}

func (s GetTextSparseEmbeddingResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s GetTextSparseEmbeddingResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponseBodyUsage) SetTokenCount(v int64) *GetTextSparseEmbeddingResponseBodyUsage {
	s.TokenCount = &v
	return s
}

type GetTextSparseEmbeddingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextSparseEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextSparseEmbeddingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextSparseEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponse) SetHeaders(v map[string]*string) *GetTextSparseEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *GetTextSparseEmbeddingResponse) SetStatusCode(v int32) *GetTextSparseEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextSparseEmbeddingResponse) SetBody(v *GetTextSparseEmbeddingResponseBody) *GetTextSparseEmbeddingResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	return nil
}

// Summary:
//
// 创建异步提取任务
//
// @param request - CreateDocumentAnalyzeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocumentAnalyzeTaskResponse
func (client *Client) CreateDocumentAnalyzeTaskWithOptions(workspaceName *string, serviceId *string, request *CreateDocumentAnalyzeTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDocumentAnalyzeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Document)) {
		body["document"] = request.Document
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["output"] = request.Output
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDocumentAnalyzeTask"),
		Version:     tea.String("2024-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v3/openapi/workspaces/" + tea.StringValue(workspaceName) + "/document-analyze/" + tea.StringValue(serviceId) + "/async"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDocumentAnalyzeTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建异步提取任务
//
// @param request - CreateDocumentAnalyzeTaskRequest
//
// @return CreateDocumentAnalyzeTaskResponse
func (client *Client) CreateDocumentAnalyzeTask(workspaceName *string, serviceId *string, request *CreateDocumentAnalyzeTaskRequest) (_result *CreateDocumentAnalyzeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDocumentAnalyzeTaskResponse{}
	_body, _err := client.CreateDocumentAnalyzeTaskWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步提取任务状态
//
// @param request - GetDocumentAnalyzeTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentAnalyzeTaskStatusResponse
func (client *Client) GetDocumentAnalyzeTaskStatusWithOptions(workspaceName *string, serviceId *string, request *GetDocumentAnalyzeTaskStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentAnalyzeTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["task_id"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentAnalyzeTaskStatus"),
		Version:     tea.String("2024-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v3/openapi/workspaces/" + tea.StringValue(workspaceName) + "/document-analyze/" + tea.StringValue(serviceId) + "/async/task-status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentAnalyzeTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步提取任务状态
//
// @param request - GetDocumentAnalyzeTaskStatusRequest
//
// @return GetDocumentAnalyzeTaskStatusResponse
func (client *Client) GetDocumentAnalyzeTaskStatus(workspaceName *string, serviceId *string, request *GetDocumentAnalyzeTaskStatusRequest) (_result *GetDocumentAnalyzeTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentAnalyzeTaskStatusResponse{}
	_body, _err := client.GetDocumentAnalyzeTaskStatusWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档相关性打分
//
// @param request - GetDocumentRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentRankResponse
func (client *Client) GetDocumentRankWithOptions(workspaceName *string, serviceId *string, request *GetDocumentRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Docs)) {
		body["docs"] = request.Docs
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentRank"),
		Version:     tea.String("2024-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v3/openapi/workspaces/" + tea.StringValue(workspaceName) + "/ranker/" + tea.StringValue(serviceId)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentRankResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档相关性打分
//
// @param request - GetDocumentRankRequest
//
// @return GetDocumentRankResponse
func (client *Client) GetDocumentRank(workspaceName *string, serviceId *string, request *GetDocumentRankRequest) (_result *GetDocumentRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentRankResponse{}
	_body, _err := client.GetDocumentRankWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档切片
//
// @param request - GetDocumentSplitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentSplitResponse
func (client *Client) GetDocumentSplitWithOptions(workspaceName *string, serviceId *string, request *GetDocumentSplitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentSplitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Document)) {
		body["document"] = request.Document
	}

	if !tea.BoolValue(util.IsUnset(request.Strategy)) {
		body["strategy"] = request.Strategy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentSplit"),
		Version:     tea.String("2024-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v3/openapi/workspaces/" + tea.StringValue(workspaceName) + "/document-split/" + tea.StringValue(serviceId)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentSplitResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档切片
//
// @param request - GetDocumentSplitRequest
//
// @return GetDocumentSplitResponse
func (client *Client) GetDocumentSplit(workspaceName *string, serviceId *string, request *GetDocumentSplitRequest) (_result *GetDocumentSplitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentSplitResponse{}
	_body, _err := client.GetDocumentSplitWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文本向量化
//
// @param request - GetTextEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextEmbeddingResponse
func (client *Client) GetTextEmbeddingWithOptions(workspaceName *string, serviceId *string, request *GetTextEmbeddingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTextEmbeddingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.InputType)) {
		body["input_type"] = request.InputType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTextEmbedding"),
		Version:     tea.String("2024-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v3/openapi/workspaces/" + tea.StringValue(workspaceName) + "/text-embedding/" + tea.StringValue(serviceId)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextEmbeddingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本向量化
//
// @param request - GetTextEmbeddingRequest
//
// @return GetTextEmbeddingResponse
func (client *Client) GetTextEmbedding(workspaceName *string, serviceId *string, request *GetTextEmbeddingRequest) (_result *GetTextEmbeddingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextEmbeddingResponse{}
	_body, _err := client.GetTextEmbeddingWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 大模型问答
//
// @param request - GetTextGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextGenerationResponse
func (client *Client) GetTextGenerationWithOptions(workspaceName *string, serviceId *string, request *GetTextGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTextGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CsiLevel)) {
		body["csi_level"] = request.CsiLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Messages)) {
		body["messages"] = request.Messages
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTextGeneration"),
		Version:     tea.String("2024-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v3/openapi/workspaces/" + tea.StringValue(workspaceName) + "/text-generation/" + tea.StringValue(serviceId)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextGenerationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大模型问答
//
// @param request - GetTextGenerationRequest
//
// @return GetTextGenerationResponse
func (client *Client) GetTextGeneration(workspaceName *string, serviceId *string, request *GetTextGenerationRequest) (_result *GetTextGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextGenerationResponse{}
	_body, _err := client.GetTextGenerationWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文本稀疏向量化
//
// @param request - GetTextSparseEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextSparseEmbeddingResponse
func (client *Client) GetTextSparseEmbeddingWithOptions(workspaceName *string, serviceId *string, request *GetTextSparseEmbeddingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTextSparseEmbeddingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.InputType)) {
		body["input_type"] = request.InputType
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnToken)) {
		body["return_token"] = request.ReturnToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTextSparseEmbedding"),
		Version:     tea.String("2024-05-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v3/openapi/workspaces/" + tea.StringValue(workspaceName) + "/text-sparse-embedding/" + tea.StringValue(serviceId)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextSparseEmbeddingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本稀疏向量化
//
// @param request - GetTextSparseEmbeddingRequest
//
// @return GetTextSparseEmbeddingResponse
func (client *Client) GetTextSparseEmbedding(workspaceName *string, serviceId *string, request *GetTextSparseEmbeddingRequest) (_result *GetTextSparseEmbeddingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextSparseEmbeddingResponse{}
	_body, _err := client.GetTextSparseEmbeddingWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}