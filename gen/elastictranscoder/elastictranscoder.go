// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package elastictranscoder provides a client for Amazon Elastic Transcoder.
package elastictranscoder

import (
	"net/http"
	"time"

	"github.com/hashicorp/aws-sdk-go/aws"
	"github.com/hashicorp/aws-sdk-go/gen/endpoints"
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// ElasticTranscoder is a client for Amazon Elastic Transcoder.
type ElasticTranscoder struct {
	client *aws.RestClient
}

// New returns a new ElasticTranscoder client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *ElasticTranscoder {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("elastictranscoder", region)

	return &ElasticTranscoder{
		client: &aws.RestClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2012-09-25",
		},
	}
}

// CancelJob the CancelJob operation cancels an unfinished job. You can
// only cancel a job that has a status of Submitted . To prevent a pipeline
// from starting to process a job while you're getting the job identifier,
// use UpdatePipelineStatus to temporarily pause the pipeline.
func (c *ElasticTranscoder) CancelJob(req *CancelJobRequest) (resp *CancelJobResponse, err error) {
	resp = &CancelJobResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobs/{Id}"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// CreateJob when you create a job, Elastic Transcoder returns data that
// includes the values that you specified plus information about the job
// that is created. If you have specified more than one output for your
// jobs (for example, one output for the Kindle Fire and another output for
// the Apple iPhone 4s), you currently must use the Elastic Transcoder API
// to list the jobs (as opposed to the AWS Console).
func (c *ElasticTranscoder) CreateJob(req *CreateJobRequest) (resp *CreateJobResponse, err error) {
	resp = &CreateJobResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobs"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// CreatePipeline the CreatePipeline operation creates a pipeline with
// settings that you specify.
func (c *ElasticTranscoder) CreatePipeline(req *CreatePipelineRequest) (resp *CreatePipelineResponse, err error) {
	resp = &CreatePipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// CreatePreset the CreatePreset operation creates a preset with settings
// that you specify. Elastic Transcoder checks the CreatePreset settings to
// ensure that they meet Elastic Transcoder requirements and to determine
// whether they comply with H.264 standards. If your settings are not valid
// for Elastic Transcoder, Elastic Transcoder returns an 400 response
// ValidationException ) and does not create the preset. If the settings
// are valid for Elastic Transcoder but aren't strictly compliant with the
// H.264 standard, Elastic Transcoder creates the preset and returns a
// warning message in the response. This helps you determine whether your
// settings comply with the H.264 standard while giving you greater
// flexibility with respect to the video that Elastic Transcoder produces.
// Elastic Transcoder uses the H.264 video-compression format. For more
// information, see the International Telecommunication Union publication
// Recommendation H.264: Advanced video coding for generic audiovisual
// services
func (c *ElasticTranscoder) CreatePreset(req *CreatePresetRequest) (resp *CreatePresetResponse, err error) {
	resp = &CreatePresetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// DeletePipeline the DeletePipeline operation removes a pipeline. You can
// only delete a pipeline that has never been used or that is not currently
// in use (doesn't contain any active jobs). If the pipeline is currently
// in use, DeletePipeline returns an error.
func (c *ElasticTranscoder) DeletePipeline(req *DeletePipelineRequest) (resp *DeletePipelineResponse, err error) {
	resp = &DeletePipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// DeletePreset the DeletePreset operation removes a preset that you've
// added in an AWS region. You can't delete the default presets that are
// included with Elastic Transcoder.
func (c *ElasticTranscoder) DeletePreset(req *DeletePresetRequest) (resp *DeletePresetResponse, err error) {
	resp = &DeletePresetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets/{Id}"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ListJobsByPipeline the ListJobsByPipeline operation gets a list of the
// jobs currently in a pipeline. Elastic Transcoder returns all of the jobs
// currently in the specified pipeline. The response body contains one
// element for each job that satisfies the search criteria.
func (c *ElasticTranscoder) ListJobsByPipeline(req *ListJobsByPipelineRequest) (resp *ListJobsByPipelineResponse, err error) {
	resp = &ListJobsByPipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobsByPipeline/{PipelineId}"

	if req.PipelineID != nil {
		uri = strings.Replace(uri, "{"+"PipelineId"+"}", aws.EscapePath(*req.PipelineID), -1)
		uri = strings.Replace(uri, "{"+"PipelineId+"+"}", aws.EscapePath(*req.PipelineID), -1)
	}

	q := url.Values{}

	if req.Ascending != nil {
		q.Set("Ascending", *req.Ascending)
	}

	if req.PageToken != nil {
		q.Set("PageToken", *req.PageToken)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ListJobsByStatus the ListJobsByStatus operation gets a list of jobs that
// have a specified status. The response body contains one element for each
// job that satisfies the search criteria.
func (c *ElasticTranscoder) ListJobsByStatus(req *ListJobsByStatusRequest) (resp *ListJobsByStatusResponse, err error) {
	resp = &ListJobsByStatusResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobsByStatus/{Status}"

	if req.Status != nil {
		uri = strings.Replace(uri, "{"+"Status"+"}", aws.EscapePath(*req.Status), -1)
		uri = strings.Replace(uri, "{"+"Status+"+"}", aws.EscapePath(*req.Status), -1)
	}

	q := url.Values{}

	if req.Ascending != nil {
		q.Set("Ascending", *req.Ascending)
	}

	if req.PageToken != nil {
		q.Set("PageToken", *req.PageToken)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ListPipelines the ListPipelines operation gets a list of the pipelines
// associated with the current AWS account.
func (c *ElasticTranscoder) ListPipelines(req *ListPipelinesRequest) (resp *ListPipelinesResponse, err error) {
	resp = &ListPipelinesResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines"

	q := url.Values{}

	if req.Ascending != nil {
		q.Set("Ascending", *req.Ascending)
	}

	if req.PageToken != nil {
		q.Set("PageToken", *req.PageToken)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ListPresets the ListPresets operation gets a list of the default presets
// included with Elastic Transcoder and the presets that you've added in an
// AWS region.
func (c *ElasticTranscoder) ListPresets(req *ListPresetsRequest) (resp *ListPresetsResponse, err error) {
	resp = &ListPresetsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets"

	q := url.Values{}

	if req.Ascending != nil {
		q.Set("Ascending", *req.Ascending)
	}

	if req.PageToken != nil {
		q.Set("PageToken", *req.PageToken)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ReadJob the ReadJob operation returns detailed information about a job.
func (c *ElasticTranscoder) ReadJob(req *ReadJobRequest) (resp *ReadJobResponse, err error) {
	resp = &ReadJobResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/jobs/{Id}"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ReadPipeline the ReadPipeline operation gets detailed information about
// a pipeline.
func (c *ElasticTranscoder) ReadPipeline(req *ReadPipelineRequest) (resp *ReadPipelineResponse, err error) {
	resp = &ReadPipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ReadPreset the ReadPreset operation gets detailed information about a
// preset.
func (c *ElasticTranscoder) ReadPreset(req *ReadPresetRequest) (resp *ReadPresetResponse, err error) {
	resp = &ReadPresetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/presets/{Id}"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// TestRole the TestRole operation tests the IAM role used to create the
// pipeline. The TestRole action lets you determine whether the IAM role
// you are using has sufficient permissions to let Elastic Transcoder
// perform tasks associated with the transcoding process. The action
// attempts to assume the specified IAM role, checks read access to the
// input and output buckets, and tries to send a test notification to
// Amazon SNS topics that you specify.
func (c *ElasticTranscoder) TestRole(req *TestRoleRequest) (resp *TestRoleResponse, err error) {
	resp = &TestRoleResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/roleTests"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// UpdatePipeline use the UpdatePipeline operation to update settings for a
// pipeline. When you change pipeline settings, your changes take effect
// immediately. Jobs that you have already submitted and that Elastic
// Transcoder has not started to process are affected in addition to jobs
// that you submit after you change settings.
func (c *ElasticTranscoder) UpdatePipeline(req *UpdatePipelineRequest) (resp *UpdatePipelineResponse, err error) {
	resp = &UpdatePipelineResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// UpdatePipelineNotifications with the UpdatePipelineNotifications
// operation, you can update Amazon Simple Notification Service (Amazon
// notifications for a pipeline. When you update notifications for a
// pipeline, Elastic Transcoder returns the values that you specified in
// the request.
func (c *ElasticTranscoder) UpdatePipelineNotifications(req *UpdatePipelineNotificationsRequest) (resp *UpdatePipelineNotificationsResponse, err error) {
	resp = &UpdatePipelineNotificationsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}/notifications"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// UpdatePipelineStatus the UpdatePipelineStatus operation pauses or
// reactivates a pipeline, so that the pipeline stops or restarts the
// processing of jobs. Changing the pipeline status is useful if you want
// to cancel one or more jobs. You can't cancel jobs after Elastic
// Transcoder has started processing them; if you pause the pipeline to
// which you submitted the jobs, you have more time to get the job IDs for
// the jobs that you want to cancel, and to send a CancelJob request.
func (c *ElasticTranscoder) UpdatePipelineStatus(req *UpdatePipelineStatusRequest) (resp *UpdatePipelineStatusResponse, err error) {
	resp = &UpdatePipelineStatusResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2012-09-25/pipelines/{Id}/status"

	if req.ID != nil {
		uri = strings.Replace(uri, "{"+"Id"+"}", aws.EscapePath(*req.ID), -1)
		uri = strings.Replace(uri, "{"+"Id+"+"}", aws.EscapePath(*req.ID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// Artwork is undocumented.
type Artwork struct {
	AlbumArtFormat aws.StringValue `json:"AlbumArtFormat,omitempty"`
	Encryption     *Encryption     `json:"Encryption,omitempty"`
	InputKey       aws.StringValue `json:"InputKey,omitempty"`
	MaxHeight      aws.StringValue `json:"MaxHeight,omitempty"`
	MaxWidth       aws.StringValue `json:"MaxWidth,omitempty"`
	PaddingPolicy  aws.StringValue `json:"PaddingPolicy,omitempty"`
	SizingPolicy   aws.StringValue `json:"SizingPolicy,omitempty"`
}

// AudioCodecOptions is undocumented.
type AudioCodecOptions struct {
	Profile aws.StringValue `json:"Profile,omitempty"`
}

// AudioParameters is undocumented.
type AudioParameters struct {
	BitRate      aws.StringValue    `json:"BitRate,omitempty"`
	Channels     aws.StringValue    `json:"Channels,omitempty"`
	Codec        aws.StringValue    `json:"Codec,omitempty"`
	CodecOptions *AudioCodecOptions `json:"CodecOptions,omitempty"`
	SampleRate   aws.StringValue    `json:"SampleRate,omitempty"`
}

// CancelJobRequest is undocumented.
type CancelJobRequest struct {
	ID aws.StringValue `json:"-"`
}

// CancelJobResponse is undocumented.
type CancelJobResponse struct {
}

// CaptionFormat is undocumented.
type CaptionFormat struct {
	Encryption *Encryption     `json:"Encryption,omitempty"`
	Format     aws.StringValue `json:"Format,omitempty"`
	Pattern    aws.StringValue `json:"Pattern,omitempty"`
}

// CaptionSource is undocumented.
type CaptionSource struct {
	Encryption *Encryption     `json:"Encryption,omitempty"`
	Key        aws.StringValue `json:"Key,omitempty"`
	Label      aws.StringValue `json:"Label,omitempty"`
	Language   aws.StringValue `json:"Language,omitempty"`
	TimeOffset aws.StringValue `json:"TimeOffset,omitempty"`
}

// Captions is undocumented.
type Captions struct {
	CaptionFormats []CaptionFormat `json:"CaptionFormats,omitempty"`
	CaptionSources []CaptionSource `json:"CaptionSources,omitempty"`
	MergePolicy    aws.StringValue `json:"MergePolicy,omitempty"`
}

// Clip is undocumented.
type Clip struct {
	TimeSpan *TimeSpan `json:"TimeSpan,omitempty"`
}

// CreateJobOutput is undocumented.
type CreateJobOutput struct {
	AlbumArt            *JobAlbumArt    `json:"AlbumArt,omitempty"`
	Captions            *Captions       `json:"Captions,omitempty"`
	Composition         []Clip          `json:"Composition,omitempty"`
	Encryption          *Encryption     `json:"Encryption,omitempty"`
	Key                 aws.StringValue `json:"Key,omitempty"`
	PresetID            aws.StringValue `json:"PresetId,omitempty"`
	Rotate              aws.StringValue `json:"Rotate,omitempty"`
	SegmentDuration     aws.StringValue `json:"SegmentDuration,omitempty"`
	ThumbnailEncryption *Encryption     `json:"ThumbnailEncryption,omitempty"`
	ThumbnailPattern    aws.StringValue `json:"ThumbnailPattern,omitempty"`
	Watermarks          []JobWatermark  `json:"Watermarks,omitempty"`
}

// CreateJobPlaylist is undocumented.
type CreateJobPlaylist struct {
	Format               aws.StringValue       `json:"Format,omitempty"`
	HlsContentProtection *HlsContentProtection `json:"HlsContentProtection,omitempty"`
	Name                 aws.StringValue       `json:"Name,omitempty"`
	OutputKeys           []string              `json:"OutputKeys,omitempty"`
}

// CreateJobRequest is undocumented.
type CreateJobRequest struct {
	Input           *JobInput           `json:"Input"`
	Output          *CreateJobOutput    `json:"Output,omitempty"`
	OutputKeyPrefix aws.StringValue     `json:"OutputKeyPrefix,omitempty"`
	Outputs         []CreateJobOutput   `json:"Outputs,omitempty"`
	PipelineID      aws.StringValue     `json:"PipelineId"`
	Playlists       []CreateJobPlaylist `json:"Playlists,omitempty"`
	UserMetadata    map[string]string   `json:"UserMetadata,omitempty"`
}

// CreateJobResponse is undocumented.
type CreateJobResponse struct {
	Job *Job `json:"Job,omitempty"`
}

// CreatePipelineRequest is undocumented.
type CreatePipelineRequest struct {
	AWSKMSKeyARN    aws.StringValue       `json:"AwsKmsKeyArn,omitempty"`
	ContentConfig   *PipelineOutputConfig `json:"ContentConfig,omitempty"`
	InputBucket     aws.StringValue       `json:"InputBucket"`
	Name            aws.StringValue       `json:"Name"`
	Notifications   *Notifications        `json:"Notifications,omitempty"`
	OutputBucket    aws.StringValue       `json:"OutputBucket,omitempty"`
	Role            aws.StringValue       `json:"Role"`
	ThumbnailConfig *PipelineOutputConfig `json:"ThumbnailConfig,omitempty"`
}

// CreatePipelineResponse is undocumented.
type CreatePipelineResponse struct {
	Pipeline *Pipeline `json:"Pipeline,omitempty"`
}

// CreatePresetRequest is undocumented.
type CreatePresetRequest struct {
	Audio       *AudioParameters `json:"Audio,omitempty"`
	Container   aws.StringValue  `json:"Container"`
	Description aws.StringValue  `json:"Description,omitempty"`
	Name        aws.StringValue  `json:"Name"`
	Thumbnails  *Thumbnails      `json:"Thumbnails,omitempty"`
	Video       *VideoParameters `json:"Video,omitempty"`
}

// CreatePresetResponse is undocumented.
type CreatePresetResponse struct {
	Preset  *Preset         `json:"Preset,omitempty"`
	Warning aws.StringValue `json:"Warning,omitempty"`
}

// DeletePipelineRequest is undocumented.
type DeletePipelineRequest struct {
	ID aws.StringValue `json:"-"`
}

// DeletePipelineResponse is undocumented.
type DeletePipelineResponse struct {
}

// DeletePresetRequest is undocumented.
type DeletePresetRequest struct {
	ID aws.StringValue `json:"-"`
}

// DeletePresetResponse is undocumented.
type DeletePresetResponse struct {
}

// Encryption is undocumented.
type Encryption struct {
	InitializationVector aws.StringValue `json:"InitializationVector,omitempty"`
	Key                  aws.StringValue `json:"Key,omitempty"`
	KeyMD5               aws.StringValue `json:"KeyMd5,omitempty"`
	Mode                 aws.StringValue `json:"Mode,omitempty"`
}

// HlsContentProtection is undocumented.
type HlsContentProtection struct {
	InitializationVector  aws.StringValue `json:"InitializationVector,omitempty"`
	Key                   aws.StringValue `json:"Key,omitempty"`
	KeyMD5                aws.StringValue `json:"KeyMd5,omitempty"`
	KeyStoragePolicy      aws.StringValue `json:"KeyStoragePolicy,omitempty"`
	LicenseAcquisitionURL aws.StringValue `json:"LicenseAcquisitionUrl,omitempty"`
	Method                aws.StringValue `json:"Method,omitempty"`
}

// Job is undocumented.
type Job struct {
	ARN             aws.StringValue   `json:"Arn,omitempty"`
	ID              aws.StringValue   `json:"Id,omitempty"`
	Input           *JobInput         `json:"Input,omitempty"`
	Output          *JobOutput        `json:"Output,omitempty"`
	OutputKeyPrefix aws.StringValue   `json:"OutputKeyPrefix,omitempty"`
	Outputs         []JobOutput       `json:"Outputs,omitempty"`
	PipelineID      aws.StringValue   `json:"PipelineId,omitempty"`
	Playlists       []Playlist        `json:"Playlists,omitempty"`
	Status          aws.StringValue   `json:"Status,omitempty"`
	UserMetadata    map[string]string `json:"UserMetadata,omitempty"`
}

// JobAlbumArt is undocumented.
type JobAlbumArt struct {
	Artwork     []Artwork       `json:"Artwork,omitempty"`
	MergePolicy aws.StringValue `json:"MergePolicy,omitempty"`
}

// JobInput is undocumented.
type JobInput struct {
	AspectRatio aws.StringValue `json:"AspectRatio,omitempty"`
	Container   aws.StringValue `json:"Container,omitempty"`
	Encryption  *Encryption     `json:"Encryption,omitempty"`
	FrameRate   aws.StringValue `json:"FrameRate,omitempty"`
	Interlaced  aws.StringValue `json:"Interlaced,omitempty"`
	Key         aws.StringValue `json:"Key,omitempty"`
	Resolution  aws.StringValue `json:"Resolution,omitempty"`
}

// JobOutput is undocumented.
type JobOutput struct {
	AlbumArt            *JobAlbumArt     `json:"AlbumArt,omitempty"`
	Captions            *Captions        `json:"Captions,omitempty"`
	Composition         []Clip           `json:"Composition,omitempty"`
	Duration            aws.LongValue    `json:"Duration,omitempty"`
	Encryption          *Encryption      `json:"Encryption,omitempty"`
	Height              aws.IntegerValue `json:"Height,omitempty"`
	ID                  aws.StringValue  `json:"Id,omitempty"`
	Key                 aws.StringValue  `json:"Key,omitempty"`
	PresetID            aws.StringValue  `json:"PresetId,omitempty"`
	Rotate              aws.StringValue  `json:"Rotate,omitempty"`
	SegmentDuration     aws.StringValue  `json:"SegmentDuration,omitempty"`
	Status              aws.StringValue  `json:"Status,omitempty"`
	StatusDetail        aws.StringValue  `json:"StatusDetail,omitempty"`
	ThumbnailEncryption *Encryption      `json:"ThumbnailEncryption,omitempty"`
	ThumbnailPattern    aws.StringValue  `json:"ThumbnailPattern,omitempty"`
	Watermarks          []JobWatermark   `json:"Watermarks,omitempty"`
	Width               aws.IntegerValue `json:"Width,omitempty"`
}

// JobWatermark is undocumented.
type JobWatermark struct {
	Encryption        *Encryption     `json:"Encryption,omitempty"`
	InputKey          aws.StringValue `json:"InputKey,omitempty"`
	PresetWatermarkID aws.StringValue `json:"PresetWatermarkId,omitempty"`
}

// ListJobsByPipelineRequest is undocumented.
type ListJobsByPipelineRequest struct {
	Ascending  aws.StringValue `json:"-"`
	PageToken  aws.StringValue `json:"-"`
	PipelineID aws.StringValue `json:"-"`
}

// ListJobsByPipelineResponse is undocumented.
type ListJobsByPipelineResponse struct {
	Jobs          []Job           `json:"Jobs,omitempty"`
	NextPageToken aws.StringValue `json:"NextPageToken,omitempty"`
}

// ListJobsByStatusRequest is undocumented.
type ListJobsByStatusRequest struct {
	Ascending aws.StringValue `json:"-"`
	PageToken aws.StringValue `json:"-"`
	Status    aws.StringValue `json:"-"`
}

// ListJobsByStatusResponse is undocumented.
type ListJobsByStatusResponse struct {
	Jobs          []Job           `json:"Jobs,omitempty"`
	NextPageToken aws.StringValue `json:"NextPageToken,omitempty"`
}

// ListPipelinesRequest is undocumented.
type ListPipelinesRequest struct {
	Ascending aws.StringValue `json:"-"`
	PageToken aws.StringValue `json:"-"`
}

// ListPipelinesResponse is undocumented.
type ListPipelinesResponse struct {
	NextPageToken aws.StringValue `json:"NextPageToken,omitempty"`
	Pipelines     []Pipeline      `json:"Pipelines,omitempty"`
}

// ListPresetsRequest is undocumented.
type ListPresetsRequest struct {
	Ascending aws.StringValue `json:"-"`
	PageToken aws.StringValue `json:"-"`
}

// ListPresetsResponse is undocumented.
type ListPresetsResponse struct {
	NextPageToken aws.StringValue `json:"NextPageToken,omitempty"`
	Presets       []Preset        `json:"Presets,omitempty"`
}

// Notifications is undocumented.
type Notifications struct {
	Completed   aws.StringValue `json:"Completed,omitempty"`
	Error       aws.StringValue `json:"Error,omitempty"`
	Progressing aws.StringValue `json:"Progressing,omitempty"`
	Warning     aws.StringValue `json:"Warning,omitempty"`
}

// Permission is undocumented.
type Permission struct {
	Access      []string        `json:"Access,omitempty"`
	Grantee     aws.StringValue `json:"Grantee,omitempty"`
	GranteeType aws.StringValue `json:"GranteeType,omitempty"`
}

// Pipeline is undocumented.
type Pipeline struct {
	ARN             aws.StringValue       `json:"Arn,omitempty"`
	AWSKMSKeyARN    aws.StringValue       `json:"AwsKmsKeyArn,omitempty"`
	ContentConfig   *PipelineOutputConfig `json:"ContentConfig,omitempty"`
	ID              aws.StringValue       `json:"Id,omitempty"`
	InputBucket     aws.StringValue       `json:"InputBucket,omitempty"`
	Name            aws.StringValue       `json:"Name,omitempty"`
	Notifications   *Notifications        `json:"Notifications,omitempty"`
	OutputBucket    aws.StringValue       `json:"OutputBucket,omitempty"`
	Role            aws.StringValue       `json:"Role,omitempty"`
	Status          aws.StringValue       `json:"Status,omitempty"`
	ThumbnailConfig *PipelineOutputConfig `json:"ThumbnailConfig,omitempty"`
}

// PipelineOutputConfig is undocumented.
type PipelineOutputConfig struct {
	Bucket       aws.StringValue `json:"Bucket,omitempty"`
	Permissions  []Permission    `json:"Permissions,omitempty"`
	StorageClass aws.StringValue `json:"StorageClass,omitempty"`
}

// Playlist is undocumented.
type Playlist struct {
	Format               aws.StringValue       `json:"Format,omitempty"`
	HlsContentProtection *HlsContentProtection `json:"HlsContentProtection,omitempty"`
	Name                 aws.StringValue       `json:"Name,omitempty"`
	OutputKeys           []string              `json:"OutputKeys,omitempty"`
	Status               aws.StringValue       `json:"Status,omitempty"`
	StatusDetail         aws.StringValue       `json:"StatusDetail,omitempty"`
}

// Preset is undocumented.
type Preset struct {
	ARN         aws.StringValue  `json:"Arn,omitempty"`
	Audio       *AudioParameters `json:"Audio,omitempty"`
	Container   aws.StringValue  `json:"Container,omitempty"`
	Description aws.StringValue  `json:"Description,omitempty"`
	ID          aws.StringValue  `json:"Id,omitempty"`
	Name        aws.StringValue  `json:"Name,omitempty"`
	Thumbnails  *Thumbnails      `json:"Thumbnails,omitempty"`
	Type        aws.StringValue  `json:"Type,omitempty"`
	Video       *VideoParameters `json:"Video,omitempty"`
}

// PresetWatermark is undocumented.
type PresetWatermark struct {
	HorizontalAlign  aws.StringValue `json:"HorizontalAlign,omitempty"`
	HorizontalOffset aws.StringValue `json:"HorizontalOffset,omitempty"`
	ID               aws.StringValue `json:"Id,omitempty"`
	MaxHeight        aws.StringValue `json:"MaxHeight,omitempty"`
	MaxWidth         aws.StringValue `json:"MaxWidth,omitempty"`
	Opacity          aws.StringValue `json:"Opacity,omitempty"`
	SizingPolicy     aws.StringValue `json:"SizingPolicy,omitempty"`
	Target           aws.StringValue `json:"Target,omitempty"`
	VerticalAlign    aws.StringValue `json:"VerticalAlign,omitempty"`
	VerticalOffset   aws.StringValue `json:"VerticalOffset,omitempty"`
}

// ReadJobRequest is undocumented.
type ReadJobRequest struct {
	ID aws.StringValue `json:"-"`
}

// ReadJobResponse is undocumented.
type ReadJobResponse struct {
	Job *Job `json:"Job,omitempty"`
}

// ReadPipelineRequest is undocumented.
type ReadPipelineRequest struct {
	ID aws.StringValue `json:"-"`
}

// ReadPipelineResponse is undocumented.
type ReadPipelineResponse struct {
	Pipeline *Pipeline `json:"Pipeline,omitempty"`
}

// ReadPresetRequest is undocumented.
type ReadPresetRequest struct {
	ID aws.StringValue `json:"-"`
}

// ReadPresetResponse is undocumented.
type ReadPresetResponse struct {
	Preset *Preset `json:"Preset,omitempty"`
}

// TestRoleRequest is undocumented.
type TestRoleRequest struct {
	InputBucket  aws.StringValue `json:"InputBucket"`
	OutputBucket aws.StringValue `json:"OutputBucket"`
	Role         aws.StringValue `json:"Role"`
	Topics       []string        `json:"Topics"`
}

// TestRoleResponse is undocumented.
type TestRoleResponse struct {
	Messages []string        `json:"Messages,omitempty"`
	Success  aws.StringValue `json:"Success,omitempty"`
}

// Thumbnails is undocumented.
type Thumbnails struct {
	AspectRatio   aws.StringValue `json:"AspectRatio,omitempty"`
	Format        aws.StringValue `json:"Format,omitempty"`
	Interval      aws.StringValue `json:"Interval,omitempty"`
	MaxHeight     aws.StringValue `json:"MaxHeight,omitempty"`
	MaxWidth      aws.StringValue `json:"MaxWidth,omitempty"`
	PaddingPolicy aws.StringValue `json:"PaddingPolicy,omitempty"`
	Resolution    aws.StringValue `json:"Resolution,omitempty"`
	SizingPolicy  aws.StringValue `json:"SizingPolicy,omitempty"`
}

// TimeSpan is undocumented.
type TimeSpan struct {
	Duration  aws.StringValue `json:"Duration,omitempty"`
	StartTime aws.StringValue `json:"StartTime,omitempty"`
}

// UpdatePipelineNotificationsRequest is undocumented.
type UpdatePipelineNotificationsRequest struct {
	ID            aws.StringValue `json:"-"`
	Notifications *Notifications  `json:"Notifications"`
}

// UpdatePipelineNotificationsResponse is undocumented.
type UpdatePipelineNotificationsResponse struct {
	Pipeline *Pipeline `json:"Pipeline,omitempty"`
}

// UpdatePipelineRequest is undocumented.
type UpdatePipelineRequest struct {
	AWSKMSKeyARN    aws.StringValue       `json:"AwsKmsKeyArn,omitempty"`
	ContentConfig   *PipelineOutputConfig `json:"ContentConfig,omitempty"`
	ID              aws.StringValue       `json:"-"`
	InputBucket     aws.StringValue       `json:"InputBucket,omitempty"`
	Name            aws.StringValue       `json:"Name,omitempty"`
	Notifications   *Notifications        `json:"Notifications,omitempty"`
	Role            aws.StringValue       `json:"Role,omitempty"`
	ThumbnailConfig *PipelineOutputConfig `json:"ThumbnailConfig,omitempty"`
}

// UpdatePipelineResponse is undocumented.
type UpdatePipelineResponse struct {
	Pipeline *Pipeline `json:"Pipeline,omitempty"`
}

// UpdatePipelineStatusRequest is undocumented.
type UpdatePipelineStatusRequest struct {
	ID     aws.StringValue `json:"-"`
	Status aws.StringValue `json:"Status"`
}

// UpdatePipelineStatusResponse is undocumented.
type UpdatePipelineStatusResponse struct {
	Pipeline *Pipeline `json:"Pipeline,omitempty"`
}

// VideoParameters is undocumented.
type VideoParameters struct {
	AspectRatio        aws.StringValue   `json:"AspectRatio,omitempty"`
	BitRate            aws.StringValue   `json:"BitRate,omitempty"`
	Codec              aws.StringValue   `json:"Codec,omitempty"`
	CodecOptions       map[string]string `json:"CodecOptions,omitempty"`
	DisplayAspectRatio aws.StringValue   `json:"DisplayAspectRatio,omitempty"`
	FixedGOP           aws.StringValue   `json:"FixedGOP,omitempty"`
	FrameRate          aws.StringValue   `json:"FrameRate,omitempty"`
	KeyframesMaxDist   aws.StringValue   `json:"KeyframesMaxDist,omitempty"`
	MaxFrameRate       aws.StringValue   `json:"MaxFrameRate,omitempty"`
	MaxHeight          aws.StringValue   `json:"MaxHeight,omitempty"`
	MaxWidth           aws.StringValue   `json:"MaxWidth,omitempty"`
	PaddingPolicy      aws.StringValue   `json:"PaddingPolicy,omitempty"`
	Resolution         aws.StringValue   `json:"Resolution,omitempty"`
	SizingPolicy       aws.StringValue   `json:"SizingPolicy,omitempty"`
	Watermarks         []PresetWatermark `json:"Watermarks,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time

var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
var _ json.RawMessage
